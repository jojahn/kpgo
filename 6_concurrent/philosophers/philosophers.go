package philosophers

import (
	"time"
	"fmt"
)

type Fork bool

type Table struct {
	forks []Fork
	count int
	putForksChannel chan int
	requestForksChannel chan int
	successChannel []chan bool 
}

func NewTable(count int) Table {
	forks := make([]Fork, count)
	for i := 0; i < count; i++ {
		forks = append(forks, true)
	}
	putForksChannel := make(chan int)
	requestForksChannel := make(chan int)
	successChannel := make([]chan bool, count)
	table := Table{
		forks: forks,
		count: count,
		putForksChannel: putForksChannel,
		requestForksChannel: requestForksChannel,
		successChannel: successChannel,
	}
	return table
}

func (t *Table) run() {
	for {
		select {
		case id := <- t.putForksChannel:
			t.forks[id] = true
			t.forks[(id + 1) % t.count] = true
			t.successChannel[id] <- true
		case id := <- t.requestForksChannel:
			if t.forks[id] && t.forks[(id + 1) % t.count] {
				t.forks[id] = false
				t.forks[(id + 1) % t.count] = false
				t.successChannel[id] <- true
			} else {
				t.successChannel[id] <- false
			}
		}
	}
}

func (x Philosopher) isHungrierThan(y Philosopher) bool {
	tx := x.lastTimeEaten.Unix() - time.Now().Unix()
	ty := y.lastTimeEaten.Unix() - time.Now().Unix()
	return tx < ty
}

type Philosopher struct {
	id int
	table *Table
	lastTimeEaten time.Time
}

func NewPhilosopher(id int, table *Table) Philosopher {
	philosopher := Philosopher{id: id, table: table}
	return philosopher
}

func (p *Philosopher) run() {
	for {
		p.takeForks() 
		p.eat()
		p.putForks()
		p.think()
	}
}

func (p *Philosopher) takeForks() {
	fmt.Printf("#%d taking forks\n", p.id)
	success := true
	for !success {
		p.table.requestForksChannel <- p.id
		success = <- p.table.successChannel[p.id]
		fmt.Println(success)
	}
}

func (p *Philosopher) eat() {
	fmt.Printf("#%d eating\n", p.id)
	time.Sleep(1 * time.Millisecond)
	p.lastTimeEaten = time.Now()
}

func (p *Philosopher) putForks() {
	fmt.Printf("#%d putting forks\n", p.id)
	p.table.putForksChannel <- p.id
}

func (p *Philosopher) think() {
	fmt.Printf("#%d thinking\n", p.id)
	time.Sleep(1 * time.Millisecond)
}