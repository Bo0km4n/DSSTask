package main

import (
	"fmt"
	"net"

	"github.com/Bo0km4n/DSSTask/architecture/layer1"
	"github.com/Bo0km4n/DSSTask/architecture/layer2"
	"github.com/Bo0km4n/DSSTask/architecture/layer3"
	"github.com/Bo0km4n/DSSTask/architecture/lib"
)

func main() {
	fmt.Println("main")
	openTCPServer()
}

func openTCPServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Listen error: %s\n", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept error: %s\n", err)
			return
		}
		defer conn.Close()

		buf := make([]byte, lib.MaxMessageLen)

		for {
			n, err := conn.Read(buf)
			if n == 0 {
				break
			}
			if err != nil {
				fmt.Printf("Read error: %s\n", err)
			}
			m := lib.NewMessage(&layer1.DIP{}, &layer2.DTCP{}, &layer2.DUDP{}, &layer3.Data{})
			m.Deserialize(buf)
			m.PP(lib.ReceiverMode)
		}
	}
}
