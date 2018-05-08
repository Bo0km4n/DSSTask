package layer3

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"

	"github.com/Bo0km4n/DSSTask/architecture/common"
)

// ValueLen //
const ValueLen = common.BufSize

// Data //
type Data struct {
	Value []byte
}

// NewData //
func NewData(v []byte) *Data {
	var dataLen int

	if len(v) >= ValueLen {
		dataLen = ValueLen
	} else {
		dataLen = len(v)
	}
	return &Data{
		Value: v[0:dataLen],
	}
}

// GetDataProtocolType //
func GetDataProtocolType() []byte {
	return []byte{0x00, 0x00, 0x00, 0x6f}
}

// GetPayloadLen //
func (d *Data) GetPayloadLen() []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(len(d.Value)))
	return bs
}

// GetPayloadMD5 //
func (d *Data) GetPayloadMD5() []byte {
	v := md5.Sum(d.Value)
	return v[:]
}

// PP //
func (d *Data) PP() {
	fmt.Printf("size = %d\n", len(d.Value))
	fmt.Println("----- layer 3 -----")

	for i := range d.Value {
		fmt.Printf("%02x ", d.Value[i])
		if (i+1)%16 == 0 {
			fmt.Printf("\n")
		}
	}
}
