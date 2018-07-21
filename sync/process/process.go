package process

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/Bo0km4n/DSSTask/sync/request"
)

type Process struct {
	ID          int
	ProcessList []*Process
	Tick        int
	reqQueue    []*request.Request
	ackQueue    []*request.Ack
	reqChan     chan *request.Request
	ackChan     chan *request.Ack
	indent      string
}

func NewProcess(id int) *Process {
	p := &Process{
		ID:          id,
		Tick:        id / 50,
		ProcessList: make([]*Process, 0),
		reqChan:     make(chan *request.Request, 10),
		ackChan:     make(chan *request.Ack, 10),
		indent: func(num int) string {
			tab := ""
			for i := 0; i < (num-1)*7; i++ {
				tab += "\t"
			}
			return tab
		}(id),
	}
	p.AddProcess(p)
	return p
}

func (p *Process) GetTick() int {
	return p.Tick
}

func (p *Process) Launch() {
	defer p.Destroy()
	tmpAckQueue := make([]*request.Ack, 0)
	var mu sync.Mutex
	go func() {
		for {
			mu.Lock()
			p.Tick++
			mu.Unlock()
		}
	}()

wait_req:
	for {
		select {
		case req := <-p.reqChan:
			if req.Tick > p.Tick {
				p.Tick = req.Tick + 1
			}
			p.reqQueue = append(p.reqQueue, req)
			a := &request.Ack{
				Request:   req,
				ProcessID: p.ID,
				Tick:      p.Tick,
			}
			tmpAckQueue = append(tmpAckQueue, a)
		case <-time.After(2 * time.Second):
			break wait_req
		}
	}

	// sort req queue
	p.SortRequest()

	// send ack list
	for _, to := range p.ProcessList {
		for i := range tmpAckQueue {
			if to.ID == p.ID {
				p.ackQueue = append(p.ackQueue, tmpAckQueue[i])
				continue
			}
			p.SendAck(to, tmpAckQueue[i])
		}
	}

wait_ack:
	for {
		select {
		case ack := <-p.ackChan:
			p.ackQueue = append(p.ackQueue, ack)
			if p.isReceivedAllAck() {
				break wait_ack
			}
		case <-time.After(5 * time.Second):
			break wait_ack
		}
	}
}

func (p *Process) AddRequest(req *request.Request) {
	p.reqChan <- req
}

func (p *Process) SendAck(to *Process, ack *request.Ack) {
	to.ackChan <- ack
}

func (p *Process) AddProcess(proc *Process) {
	p.ProcessList = append(p.ProcessList, proc)
}

func (p *Process) SortRequest() {
	sort.Slice(p.reqQueue, func(i, j int) bool {
		return p.reqQueue[i].Tick < p.reqQueue[j].Tick
	})
}

func (p *Process) isReceivedAllAck() bool {
	return len(p.ackQueue) == len(p.ProcessList)*len(p.reqQueue)
}

func (p *Process) Dump() {
	sort.Slice(p.ackQueue, func(i, j int) bool {
		return p.ackQueue[i].Tick < p.ackQueue[j].Tick
	})

	table := map[string]int{}

	for i := range p.ackQueue {
		table[p.ackQueue[i].Dump()] = p.ackQueue[i].Tick
	}
	for i := range p.reqQueue {
		table[p.reqQueue[i].Dump()] = p.reqQueue[i].Tick
	}

	// sort map
	a := list{}
	for k, v := range table {
		e := entry{k, v}
		a = append(a, e)
	}

	sort.Sort(a)
	for i := range a {
		p.print(a[i].key)
	}
}

type entry struct {
	key   string
	value int
}
type list []entry

func (l list) Len() int {
	return len(l)
}

func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l list) Less(i, j int) bool {
	if l[i].value == l[j].value {
		return (l[i].key < l[j].key)
	} else {
		return (l[i].value < l[j].value)
	}
}

func (p *Process) print(arg interface{}) {
	fmt.Printf("%s%v\n", p.indent, arg)
}

func (p *Process) Destroy() {
	close(p.ackChan)
	close(p.reqChan)
}
