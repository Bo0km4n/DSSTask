package layer2

import (
	"bytes"
	"fmt"
	"log"

	"github.com/Bo0km4n/DSSTask/architecture/common"
)

// DTCP //
type DTCP struct {
	Type   []byte
	Len    []byte
	Digest []byte
}

// NewDTCP //
func NewDTCP(t, len, digest []byte) *DTCP {
	return &DTCP{
		Type:   t[0:4],
		Len:    len[0:4],
		Digest: digest[0:16],
	}
}

// IsTCPProtocol //
func IsTCPProtocol(b []byte) bool {
	t := common.GetTCPProtocolType()
	for i := range t {
		if b[i] != t[i] {
			return false
		}
	}
	return true
}

// PP //
func (d *DTCP) PP(mode int, digest []byte) {
	t := common.ConvertBytesToInt(d.Type)
	len := common.ConvertBytesToInt(d.Len)
	fmt.Println("----- layer2 -----")
	fmt.Printf("%-10.10s = %d\n", "type", t)
	fmt.Printf("%-10.10s = %d\n", "len", len)

	for i := range d.Digest {
		fmt.Printf("%02x ", d.Digest[i])
		if (i+1)%16 == 0 {
			fmt.Printf("\n")
		}
	}

	if mode == 1 {
		for i := range digest {
			fmt.Printf("%02x ", digest[i])
			if (i+1)%16 == 0 {
				fmt.Printf("\n")
			}
		}

		if bytes.Compare(digest, d.Digest) == 0 {
			fmt.Println("packet is protected")
		} else {
			log.Fatal("[FATAL] packet is tempered")
		}
	}
}
