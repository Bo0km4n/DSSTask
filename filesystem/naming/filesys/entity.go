package filesys

type FileSys struct {
	SIsize  uint16
	SFsize  uint16
	SNfree  uint16
	Sfree   [100]uint16
	SNInode uint16
	SInode  [100]uint16
	SFlock  byte
	SIlock  byte
	SFmod   byte
	SRonly  byte
	STime   [2]uint16
	Pad     [50]uint16
}
