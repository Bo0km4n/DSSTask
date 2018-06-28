package disk

import (
	gobytes "bytes"
	"encoding/binary"
	"log"
	"os"

	"github.com/Bo0km4n/DSSTask/filesystem/naming/byte"
	"github.com/Bo0km4n/DSSTask/filesystem/naming/filesys"
	"github.com/Bo0km4n/DSSTask/filesystem/naming/inode"
)

const BLOCK = 512
const INODE_SIZE = 32
const ROOT = 1

type Disk struct {
	BootArea    bytes.BytesT
	SupreArea   bytes.BytesT
	InodeArea   bytes.BytesT
	StorageArea bytes.BytesT
	FileSys     *filesys.FileSys
	Inodes      []*inode.Inode
}

func NewDisk() *Disk {
	return &Disk{
		FileSys: &filesys.FileSys{},
		Inodes:  make([]*inode.Inode, 0),
	}
}

func (d *Disk) Load(fileName string) error {
	fp, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fp.Close()

	log.Println("Step: reading boot area... ")
	d.BootArea = readBlocks(fp, 1)
	log.Println("done!")

	log.Println("Step: reading super area... ")
	d.SupreArea = readBlocks(fp, 1)
	log.Println("done!")

	d.assignFileSys(d.SupreArea)
	d.assignInode(fp)
	d.assigneStorage(fp)
	return nil
}

func readBlocks(fp *os.File, blocks int) bytes.BytesT {
	var b bytes.BytesT
	b.Len = BLOCK * blocks
	b.Head = make([]byte, b.Len)

	c, err := fp.Read(b.Head)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("read %d bytes from file\n", c)
	return b
}

func (d *Disk) assignFileSys(block bytes.BytesT) {
	log.Println("Step: assigning file sys to struct... ")

	binary.Read(gobytes.NewBuffer(block.Head[0:2]), binary.LittleEndian, &d.FileSys.SIsize)
	binary.Read(gobytes.NewBuffer(block.Head[2:4]), binary.LittleEndian, &d.FileSys.SFsize)
	binary.Read(gobytes.NewBuffer(block.Head[4:6]), binary.LittleEndian, &d.FileSys.SNfree)
	binary.Read(gobytes.NewBuffer(block.Head[6:206]), binary.LittleEndian, &d.FileSys.Sfree)
	binary.Read(gobytes.NewBuffer(block.Head[206:208]), binary.LittleEndian, &d.FileSys.SNInode)
	binary.Read(gobytes.NewBuffer(block.Head[208:408]), binary.LittleEndian, &d.FileSys.SInode)
	binary.Read(gobytes.NewBuffer(block.Head[408:409]), binary.LittleEndian, &d.FileSys.SFlock)
	binary.Read(gobytes.NewBuffer(block.Head[409:410]), binary.LittleEndian, &d.FileSys.SIlock)
	binary.Read(gobytes.NewBuffer(block.Head[410:411]), binary.LittleEndian, &d.FileSys.SFmod)
	binary.Read(gobytes.NewBuffer(block.Head[411:412]), binary.LittleEndian, &d.FileSys.SRonly)
	binary.Read(gobytes.NewBuffer(block.Head[412:416]), binary.LittleEndian, &d.FileSys.STime)
	binary.Read(gobytes.NewBuffer(block.Head[416:512]), binary.LittleEndian, &d.FileSys.Pad)

	log.Println("done!")
}

func (d *Disk) assigneStorage(fp *os.File) {
	log.Println("Step: assigning storage area to file sys struct...")

	sfsize := int32(d.FileSys.SFsize)
	d.StorageArea = readBlocks(fp, int(sfsize))
	log.Println("done!")
}

func (d *Disk) assignInode(fp *os.File) {
	log.Println("Step: assigning inode area to inode struct... ")

	var b bytes.BytesT
	blockLen := int(int32(d.FileSys.SIsize))
	b.Len = BLOCK * blockLen
	b.Head = make([]byte, b.Len)
	c, err := fp.Read(b.Head)
	if err != nil {
		log.Fatal(err)
	}

	d.InodeArea = b

	// assigne inodes
	for i := 0; i < d.InodeArea.Len; i += INODE_SIZE {
		in := castInode(d.InodeArea.Head[i : i+INODE_SIZE])
		d.Inodes = append(d.Inodes, in)
	}
	log.Printf("read %d bytes from file\n", c)
	log.Println("done!")
	return
}

func castInode(b []byte) *inode.Inode {
	in := inode.Inode{}

	if len(b) < 32 {
		return nil
	}

	binary.Read(gobytes.NewBuffer(b[0:2]), binary.LittleEndian, &in.Imode)
	binary.Read(gobytes.NewBuffer(b[2:3]), binary.LittleEndian, &in.INlink)
	binary.Read(gobytes.NewBuffer(b[3:4]), binary.LittleEndian, &in.IUid)
	binary.Read(gobytes.NewBuffer(b[4:5]), binary.LittleEndian, &in.IGid)
	binary.Read(gobytes.NewBuffer(b[5:6]), binary.LittleEndian, &in.ISize0)
	binary.Read(gobytes.NewBuffer(b[6:8]), binary.LittleEndian, &in.ISize1)
	binary.Read(gobytes.NewBuffer(b[8:24]), binary.LittleEndian, &in.IAddr)
	binary.Read(gobytes.NewBuffer(b[24:28]), binary.LittleEndian, &in.IAttime)
	binary.Read(gobytes.NewBuffer(b[28:32]), binary.LittleEndian, &in.IMttime)
	return &in
}

func (d *Disk) GetInode(index int) *inode.Inode {
	return d.Inodes[index-1]
}
