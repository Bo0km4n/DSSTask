package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"github.com/Bo0km4n/DSSTask/filesystem/naming/disk"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Reader) {
	scanner := bufio.NewScanner(in)
	d := disk.NewDisk()
	if err := d.Load("v6root"); err != nil {
		log.Fatal(err)
	}
	wd := d.GetInode(disk.ROOT)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		fmt.Println(line)
	}

}
