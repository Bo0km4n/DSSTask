package inode

import (
	"bytes"
)

type Inode struct {
	Imode   uint16
	INlink  byte
	IUid    byte
	IGid    byte
	ISize0  byte
	ISize1  uint16
	IAddr   [8]uint16
	IAttime [2]uint16
	IMttime [2]uint16
}

const (
	IFDIR = 040000
)

func (i *Inode) GetFileSize() int {
	return int(i.ISize0<<8) + int(i.ISize1)
}

func (i *Inode) GetDetail() string {
	buf := bytes.NewBufferString("")

	if i.Imode&IFDIR != 0x00 {
		buf.WriteString("d")
	} else {
		buf.WriteString("-")
	}

	if i.Imode&0400 != 0x00 {
		buf.WriteString("r")
	} else {
		buf.WriteString("-")
	}

	if i.Imode&0200 != 0x00 {
		buf.WriteString("w")
	} else {
		buf.WriteString("-")
	}

	if i.Imode&0100 != 0x00 {
		buf.WriteString("x")
	} else {
		buf.WriteString("-")
	}

	if i.Imode&040 != 0x00 {
		buf.WriteString("r")
	} else {
		buf.WriteString("-")
	}

	if i.Imode&020 != 0x00 {
		buf.WriteString("w")
	} else {
		buf.WriteString("-")
	}

	if i.Imode&010 != 0x00 {
		buf.WriteString("x")
	} else {
		buf.WriteString("-")
	}

	if i.Imode&04 != 0x00 {
		buf.WriteString("r")
	} else {
		buf.WriteString("-")
	}

	if i.Imode&02 != 0x00 {
		buf.WriteString("w")
	} else {
		buf.WriteString("-")
	}

	if i.Imode&01 != 0x00 {
		buf.WriteString("x")
	} else {
		buf.WriteString("-")
	}

	return buf.String()
}
