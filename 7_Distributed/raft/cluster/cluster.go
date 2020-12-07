package cluster

import "github.com/jojahn/kpgo/7_Distributed/raft/node"

type Cluster struct {
	count int
	nodes                       []node.Node
	RequestVotesChannels        []chan node.RequestVoteArgs
	RequestVotesResultsChannels []chan node.RequestVoteResults
	AppendEntriesChannel        []chan node.AppendEntriesArgs
	AppendEntriesResultsChannel []chan node.AppendEntriesResults
}

func NewCluster(count int) Cluster {
	cluster := Cluster{count: count}
	for i := 0; i < count; i++ {
		cluster.addNode()
	}
	return cluster
}

func (c *Cluster) addNode() {
	id := len(c.nodes) - 1
	n := node.NewNode(c.count, id, c)
	c.nodes = append(c.nodes, n)
	c.RequestVotesChannels[id] = make(chan node.RequestVoteArgs)
	c.RequestVotesResultsChannels[id] = make(chan node.RequestVoteResults)
	c.AppendEntriesChannel[id] = make(chan node.AppendEntriesArgs)
	c.AppendEntriesResultsChannel[id] = make(chan node.AppendEntriesResults)
}