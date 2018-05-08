package layer2

import (
	"fmt"

	"github.com/Bo0km4n/DSSTask/architecture/common"
)

// DUDP //
type DUDP struct {
	Type []byte
	Len  []byte
}

// NewDUDP //
func NewDUDP(t, len []byte) *DUDP {
	return &DUDP{
		Type: t[0:4],
		Len:  len[0:4],
	}
}

// IsUDPProtocol //
func IsUDPProtocol(b []byte) bool {
	t := common.GetUDPProtocolType()
	for i := range t {
		if b[i] != t[i] {
			return false
		}
	}
	return true
}

// PP //
func (d *DUDP) PP() {
	t := common.ConvertBytesToInt(d.Type)
	len := common.ConvertBytesToInt(d.Len)
	fmt.Println("----- layer2 -----")
	fmt.Printf("%-10.10s = %d\n", "type", t)
	fmt.Printf("%-10.10s = %d\n", "len", len)
}
