package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/Bo0km4n/DSSTask/filesystem/naming/inode"

	"github.com/Bo0km4n/DSSTask/filesystem/naming/disk"
)

const PROMPT = ">> "

var wd *inode.Inode

func Start(in io.Reader, out io.Reader) {
	scanner := bufio.NewScanner(in)
	d := disk.NewDisk()
	if err := d.Load("v6root"); err != nil {
		log.Fatal(err)
	}
	wd = d.GetInode(disk.ROOT)
	d.LoadFile(wd)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		exec(line)
	}
}

func exec(stmt string) {
	args := strings.Split(stmt, " ")
	cmd := args[0]
	opts := args[1:]

	switch cmd {
	case "ls":
		fmt.Println("ls", opts)
	case "cd":
		fmt.Println("cd", opts)
	default:
		fmt.Println("not supported ", cmd)
	}
}
