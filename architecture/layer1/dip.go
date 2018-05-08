package layer1

import (
	"fmt"

	"github.com/Bo0km4n/DSSTask/architecture/common"
)

// DIP layer1 struct
type DIP struct {
	Type    []byte
	Version []byte
	TTL     []byte
}

// NewDIP return dip object
func NewDIP(t, version, ttl []byte) *DIP {
	return &DIP{
		Type:    t[0:4],
		Version: version[0:4],
		TTL:     ttl[0:4],
	}
}

// SetType //
func (d *DIP) SetType(b []byte) {
	if len(b) != 4 {
		return
	}
	d.Type = b
}

// SetVersion //
func (d *DIP) SetVersion(b []byte) {
	if len(b) != 4 {
		return
	}
	d.Version = b
}

// SetTTL //
func (d *DIP) SetTTL(b []byte) {
	if len(b) != 4 {
		return
	}
	d.TTL = b
}

// GetType //
func (d *DIP) GetType() []byte {
	return d.Type[0:4]
}

// GetVersion //
func (d *DIP) GetVersion() []byte {
	return d.Version[0:4]
}

// GetTTL //
func (d *DIP) GetTTL() []byte {
	return d.TTL[0:4]
}

// DIPVersion //
func DIPVersion() []byte {
	return []byte{0x00, 0x00, 0x00, 0x0a}
}

// PP //
func (d *DIP) PP() {
	t := common.ConvertBytesToInt(d.Type)
	version := common.ConvertBytesToInt(d.Version)
	ttl := common.ConvertBytesToInt(d.TTL)
	fmt.Println("----- layer1 -----")
	fmt.Printf("%-10.10s = %d\n", "type", t)
	fmt.Printf("%-10.10s = %d\n", "version", version)
	fmt.Printf("%-10.10s = %d\n", "ttl", ttl)
}
