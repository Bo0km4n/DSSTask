package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	// BufSize //
	BufSize = 1024
)

// AppendByte //
func AppendByte(v, s []byte) []byte {
	for i := range s {
		v = append(v, s[i])
	}
	return v
}

// GetLayer2Type //
func GetLayer2Type(t string) []byte {
	switch t {
	case "tcp":
		return GetTCPProtocolType()
	case "udp":
		return GetUDPProtocolType()
	default:
		return GetUDPProtocolType()
	}
}

// TTL //
func TTL() []byte {
	return []byte{0x00, 0x00, 0xff, 0xff}
}

// TCPVersion //
func TCPVersion() []byte {
	return []byte{0x00, 0x00, 0x00, 0x0a}
}

// ConvertBytesToInt //
func ConvertBytesToInt(b []byte) int32 {
	var i int32
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.BigEndian, &i)
	if err != nil {
		fmt.Println("binary.Read failded:", err)
	}

	return i
}

// GetTCPProtocolType //
func GetTCPProtocolType() []byte {
	return []byte{0x00, 0x00, 0x00, 0x06}
}

// GetUDPProtocolType //
func GetUDPProtocolType() []byte {
	return []byte{0x00, 0x00, 0x00, 0x11}
}
