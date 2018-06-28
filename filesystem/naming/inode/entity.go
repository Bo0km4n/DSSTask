package inode

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

func (i *Inode) GetFileSize() int {
	return int(uint16(i.ISize0)<<8 + i.ISize1)
}
