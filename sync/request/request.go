package request

import (
	"fmt"
)

type Request struct {
	ID       int
	ClientID int
	Tick     int
}

type Ack struct {
	Request   *Request
	ProcessID int
	Tick      int
}

// Dump //
func (r *Request) Dump() string {
	return fmt.Sprintf("REQ%d: %d.%d", r.ID, r.Tick, r.ClientID)
}

// Dump //
func (a *Ack) Dump() string {
	return fmt.Sprintf("ACK%d-%d: %d.%d", a.Request.ID, a.ProcessID, a.Tick, a.ProcessID)
}
