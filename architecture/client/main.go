package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/Bo0km4n/DSSTask/architecture/common"
	"github.com/Bo0km4n/DSSTask/architecture/layer1"
	"github.com/Bo0km4n/DSSTask/architecture/layer2"
	"github.com/Bo0km4n/DSSTask/architecture/layer3"
	"github.com/Bo0km4n/DSSTask/architecture/lib"
)

func main() {
	if len(os.Args) < 3 {
		return
	}

	file, _ := os.Open(os.Args[2])
	defer file.Close()

	buf := make([]byte, common.BufSize)

	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Dial error: %s\n", err)
		return
	}
	defer conn.Close()

	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		body := newPayload(os.Args[1], buf)
		sendMessage(conn, body)
	}
}

func sendMessage(conn net.Conn, m []byte) {
	conn.Write(m)
}

func newPayload(layer2Type string, fileContents []byte) []byte {
	var dudp *layer2.DUDP
	var dtcp *layer2.DTCP

	data := layer3.NewData(fileContents)
	dip := layer1.NewDIP(common.GetLayer2Type(layer2Type), layer1.DIPVersion(), common.TTL())

	if layer2.IsTCPProtocol(common.GetLayer2Type(layer2Type)) {
		dtcp = layer2.NewDTCP(layer3.GetDataProtocolType(), data.GetPayloadLen(), data.GetPayloadMD5()[0:16])
		dudp = nil
	} else {
		dudp = layer2.NewDUDP(layer3.GetDataProtocolType(), data.GetPayloadLen())
		dtcp = nil
	}

	m := lib.NewMessage(dip, dtcp, dudp, data)
	m.PP(lib.SenderMode)

	return m.Serialize()
}

func readFile(fileName string) []byte {
	contents, _ := ioutil.ReadFile(fileName)
	if len(contents) >= layer3.ValueLen {
		return contents[0:layer3.ValueLen]
	}
	return contents[0:len(contents)]
}
