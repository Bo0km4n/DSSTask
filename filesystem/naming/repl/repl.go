package repl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/Bo0km4n/DSSTask/filesystem/naming/inode"

	"github.com/Bo0km4n/DSSTask/filesystem/naming/disk"
)

// PROMPT //
const PROMPT = ">> "

var wd *inode.Inode
var d *disk.Disk

// Start starts repl interface
func Start(in io.Reader, out io.Reader) {
	scanner := bufio.NewScanner(in)
	d = disk.NewDisk()
	if err := d.Load("v6root"); err != nil {
		log.Fatal(err)
	}
	wd = d.GetInode(disk.ROOT)
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
		ls(opts...)
	case "cd":
		fmt.Println("cd", opts)
	default:
		fmt.Println("not supported ", cmd)
	}
}

func ls(args ...string) {
	dir := d.LoadFile(wd)

	entries := d.AssignBytesToEntries(dir)

	printBuffer := bytes.NewBufferString("")

	for _, v := range entries {
		if isContainArgs(args, "-l") && v.Ino != 0x0000 {
			inode := d.GetInode(int(v.Ino))
			printBuffer.WriteString(inode.GetDetail())
			printBuffer.WriteString(fmt.Sprintf("%10d ", inode.GetFileSize()))
		}

		if v.Ino != 0x0000 {
			printBuffer.WriteString(v.GetName())
			fmt.Println(printBuffer.String())
		}

		printBuffer = bytes.NewBufferString("")
	}
}

func isContainArgs(args []string, arg string) bool {
	for _, v := range args {
		if strings.Contains(v, arg) {
			return true
		}
	}
	return false
}
