package node

import (
	_ "fmt"
	"github.com/jojahn/kpgo/7_Distributed/raft/cluster"
	"math/rand"
	"time"
)

const (
	electionTimeoutMinMs = 1500                    // min. milliseconds to wait for a heartbeat (starts election when timing out)
	electionTimeoutMaxMs = 3000                    // max. milliseconds to wait for a heartbeat (starts election when timing out)
	startupWait          = 1000 * time.Millisecond // time to wait for connection to other nodes
	voteTimeout          = 1500 * time.Millisecond // time to wait for votes from other clients
	heartbeatTimeout     = 500 * time.Millisecond  // time to wait for a client to accept a heartbeat
	heartbeatTicker      = 1000 * time.Millisecond // time between two heartbeats
	dataTimeout          = 1000 * time.Millisecond // time to wait when passing requests to the leader
)

type Entry struct {
	term int
	data []byte
}

type mode int
const (
	follower mode = iota
	candidate
	leader
)

type Node struct {
	state mode
	count int
	Id int
	cluster *cluster.Cluster
	stopped bool

	currentTerm int
	votedFor *int
	log []Entry

	commitIndex int
	lastApplied int

	nextIndex []int
	matchIndex []int
}

// Initial Boot
func NewNode(count int, id int, cluster *cluster.Cluster) Node {
	node := Node{
		count: count,
		Id: id,
		cluster: cluster,
		votedFor: nil,
		currentTerm: 0,
	}
	return node
}

// Boot
func (n *Node) Start() {
	n.stopped = false
	n.state = follower
	n.run()
}

func (n *Node) Stop() {
	n.stopped = true
}

func (n *Node) run() {
	for !n.stopped {
		electionTimeout := time.Duration(
			rand.Intn(electionTimeoutMaxMs-electionTimeoutMinMs) + electionTimeoutMinMs) * time.Millisecond
		electionTimer := time.NewTimer(electionTimeout)
		timer := time.NewTimer(heartbeatTimeout)
		select {
			case <- electionTimer.C:
				n.onElectionTimeout()
			case args := <- n.cluster.RequestVotesChannels[n.id]:
				n.onReceiveVoteRequest(args)
			case <- timer.C: {
				if n.state == leader {
					n.onAppendEntries()
				}
			}
			case args := <- n.cluster.AppendEntriesChannel[n.id]:
				n.onReceiveAppendEntries(args)
		}
	}
}

func (n *Node) onReceiveVoteRequest(args RequestVoteArgs) {
	if args.term < n.currentTerm {
		n.cluster.RequestVotesResultsChannels[n.id] <- RequestVoteResults{
			term:        n.currentTerm,
			voteGranted: false,
		}
	} else {
		n.currentTerm = args.term
	}
	if *n.votedFor == args.candidateId || n.votedFor == nil {
		n.cluster.RequestVotesResultsChannels[n.id] <- RequestVoteResults{
			term:        n.currentTerm,
			voteGranted: true,
		}
		n.votedFor = &args.candidateId
	}
	n.cluster.RequestVotesResultsChannels[n.id] <- RequestVoteResults{
		term:        n.currentTerm,
		voteGranted: false,
	}
}

func (n *Node) onElectionTimeout() {
	electionTimeout := time.Duration(
		rand.Intn(electionTimeoutMaxMs-electionTimeoutMinMs) + electionTimeoutMinMs) * time.Millisecond
	electionTimer := time.NewTimer(electionTimeout)

	n.currentTerm++
	n.state = candidate
	n.cluster.RequestVotesResultsChannels[n.id] <- RequestVoteResults{
		term:        n.currentTerm,
		voteGranted: true,
	}

	channel := n.requestVotes(n.currentTerm, n.id)
	select {
	case votesGranted := <- channel:
		if votesGranted {
			n.cluster.SetLeader(*n)
			n.state = leader
			n.nextIndex = make([]int, 0)
			n.matchIndex = make([]int, 0)
		} else {
			n.state = follower
		}
	case <- electionTimer.C: {
		n.onElectionTimeout()
	}
	}
}

func (n *Node) onAppendEntries() {
	timer := time.NewTimer(heartbeatTimeout)

	channel := n.appendEntries()
	select {
	case results := <- channel:
		if results.success {

		}
	case <- timer.C: {
		n.onElectionTimeout()
	}
	}
}

func (n *Node) onReceiveAppendEntries(args AppendEntriesArgs) {

}

type AppendEntriesArgs struct {
	term int
	leaderId int
	prevLogIndex int
	prevLogTerm int
	entries []Entry
	leaderCommit int
}

type AppendEntriesResults struct {
	term int
	success bool
}

func (n *Node) appendEntries(args AppendEntriesArgs) <-chan AppendEntriesResults {
	c := n.cluster
	for _, ch := range c.RequestVotesChannels {
		ch <- RequestVoteArgs{
			term:         term,
			candidateId:  candidateId,
			lastLogIndex: 0,
			lastLogTerm:  0,
		}
	}
	channel := make(chan bool)
	votes := 0
	for _, ch := range c.RequestVotesResultsChannels {
		results := <- ch
		if results.voteGranted {
			votes++
		}
		if results.term > n.currentTerm {
			n.currentTerm = results.term
		}
	}
	go func() { channel <- votes > n.count / 2 }()
	return channel
}

type RequestVoteArgs struct {
	term int
	candidateId int
	lastLogIndex int
	lastLogTerm int
}

type RequestVoteResults struct {
	term int
	voteGranted bool
}

func (n *Node) requestVotes(term int, candidateId int) <-chan bool {
	c := n.cluster
	for _, ch := range c.RequestVotesChannels {
		ch <- RequestVoteArgs{
			term:         term,
			candidateId:  candidateId,
			lastLogIndex: 0,
			lastLogTerm:  0,
		}
	}
	channel := make(chan bool)
	votes := 0
	for _, ch := range c.RequestVotesResultsChannels {
		results := <- ch
		if results.voteGranted {
			votes++
		}
		if results.term > n.currentTerm {
			n.currentTerm = results.term
		}
	}
	go func() { channel <- votes > n.count / 2 }()
	return channel
}
/*
func (n *Node) RequestVote(term int, candidateId int, lastLogIndex int, lastLogTerm int) (int, bool) {
	c := n.cluster
	for _, ch := range c.RequestVotesChannels {
		ch <- RequestVoteArgs{
			term:         term,
			candidateId:  candidateId,
			lastLogIndex: lastLogIndex,
			lastLogTerm:  lastLogTerm,
		}
	}
	votes := 0
	for _, ch := range c.RequestVotesResultsChannels {
		results := <- ch
		if results.voteGranted {
			votes++
		}
	}
	return 0, votes > n.count / 2
}
*/