package main

import (
	"sync"
	"time"

	"github.com/Bo0km4n/DSSTask/sync/process"
	"github.com/Bo0km4n/DSSTask/sync/request"
)

func main() {
	p1 := process.NewProcess(1)
	p2 := process.NewProcess(2)
	wg := &sync.WaitGroup{}

	p1.AddProcess(p2)
	p2.AddProcess(p1)

	wg.Add(1)
	go func() {
		p1.Launch()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		p2.Launch()
		wg.Done()
	}()

	time.Sleep(time.Millisecond * 1)

	cs := []*request.Request{
		{
			ID:       1,
			ClientID: 1,
		},
		{
			ID:       2,
			ClientID: 2,
		},
	}

	for _, req := range cs {
		req.Tick = p1.GetTick()
		p1.AddRequest(req)
		req.Tick = p2.GetTick()
		p2.AddRequest(req)
	}

	wg.Wait()

	p1.Dump()
	p2.Dump()
}
