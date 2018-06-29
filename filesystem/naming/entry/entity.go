package entry

import "strings"

type Entry struct {
	Ino  uint16
	Name [14]byte
}

func (e *Entry) GetName() string {
	return strings.Replace(string(e.Name[:]), "\x00", "", -1)
}

func (e *Entry) GetIno() int {
	return int(e.Ino)
}
