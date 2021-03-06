package repl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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
		cd(opts...)
	case "exit":
		log.Println("bye")
		os.Exit(0)
	default:
		fmt.Println("not supported ", cmd)
	}
}

func ls(args ...string) {
	dir := d.LoadFile(wd)

	entries := d.AssignBytesToEntries(dir)

	// pp.Println(entries)

	printBuffer := bytes.NewBufferString("")

	for _, v := range entries {
		if isContainArgs(args, "-l") && v.Ino != 0x0000 {
			if len(d.Inodes) > int(v.Ino) {
				inode := d.GetInode(int(v.Ino))
				printBuffer.WriteString(inode.GetDetail())
				printBuffer.WriteString(fmt.Sprintf("%10d ", inode.GetFileSize()))
			}
		}

		if v.Ino != 0x0000 {
			if len(d.Inodes) > int(v.Ino) {
				printBuffer.WriteString(v.GetName())
				fmt.Println(printBuffer.String())
			}
		}

		printBuffer = bytes.NewBufferString("")
	}
}

func cd(args ...string) {
	paths := strings.Split(args[0], "/")
	node := wd
	var target *inode.Inode

	if paths[0] == "" || len(args) == 0 {
		node = d.GetInode(disk.ROOT)
	}

	for idx := range paths {
		if paths[idx] == "" {
			continue
		}
		dir := d.LoadFile(node)
		entries := d.AssignBytesToEntries(dir)
		for i := range entries {
			if paths[idx] == entries[i].GetName() {
				target = d.GetInode(int(entries[i].Ino))
			}
		}

		if target.Imode&inode.IFDIR != 0x00 {
			node = target
		} else {
			node = nil
			break
		}
	}

	if node != nil {
		wd = node
	} else {
		log.Println("no such directory")
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
