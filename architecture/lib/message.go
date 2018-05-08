package lib

import (
	"github.com/Bo0km4n/DSSTask/architecture/common"
	"github.com/Bo0km4n/DSSTask/architecture/layer1"
	"github.com/Bo0km4n/DSSTask/architecture/layer2"
	"github.com/Bo0km4n/DSSTask/architecture/layer3"
)

// MaxMessageLen = payload + dtcp + dip
const MaxMessageLen = common.BufSize + 24 + 12
const (
	// SenderMode //
	SenderMode = iota
	// ReceiverMode //
	ReceiverMode
)

// Message //
type Message struct {
	DIP  *layer1.DIP
	DTCP *layer2.DTCP
	DUDP *layer2.DUDP
	Data *layer3.Data
}

// NewMessage //
func NewMessage(dip *layer1.DIP, dtcp *layer2.DTCP, dudp *layer2.DUDP, data *layer3.Data) *Message {
	return &Message{
		DIP:  dip,
		DTCP: dtcp,
		DUDP: dudp,
		Data: data,
	}
}

// Serialize //
func (m *Message) Serialize() []byte {
	v := make([]byte, 0)

	v = common.AppendByte(v, m.DIP.Type[0:4])
	v = common.AppendByte(v, m.DIP.Version[0:4])
	v = common.AppendByte(v, m.DIP.TTL[0:4])

	if layer2.IsTCPProtocol(m.DIP.Type[0:4]) {
		v = common.AppendByte(v, m.DTCP.Type[0:4])
		v = common.AppendByte(v, m.DTCP.Len[0:4])
		v = common.AppendByte(v, m.DTCP.Digest[0:16])
	} else {
		v = common.AppendByte(v, m.DUDP.Type[0:4])
		v = common.AppendByte(v, m.DUDP.Len[0:4])
	}

	v = common.AppendByte(v, m.Data.Value)
	return v
}

// Deserialize //
func (m *Message) Deserialize(b []byte) {
	m.DIP.Type = b[0:4]
	m.DIP.Version = b[4:8]
	m.DIP.TTL = b[8:12]

	if layer2.IsTCPProtocol(m.DIP.Type[0:4]) {
		m.DTCP.Type = b[12:16]
		m.DTCP.Len = b[16:20]
		m.DTCP.Digest = b[20:36]

		m.Data.Value = b[36:len(b)]
	} else {
		m.DUDP.Type = b[12:16]
		m.DUDP.Len = b[16:20]

		m.Data.Value = b[20:len(b)]
	}
}

// PP PrettyPrinter
func (m *Message) PP(order int) {
	switch order {
	case SenderMode:
		m.ppSender()
	case ReceiverMode:
		m.ppReceiver()
	default:
		m.ppSender()
	}
}

func (m *Message) ppSender() {
	m.Data.PP()
	if m.DTCP != nil {
		m.DTCP.PP(SenderMode, nil)
	} else {
		m.DUDP.PP()
	}
	m.DIP.PP()
}

func (m *Message) ppReceiver() {
	m.DIP.PP()
	if layer2.IsTCPProtocol(m.DIP.Type[0:4]) {
		m.DTCP.PP(ReceiverMode, m.Data.GetPayloadMD5())
	} else {
		m.DUDP.PP()
	}
	m.Data.PP()
}
