package entry

type Entry struct {
	Ino  uint16
	Name [14]byte
}

func (e *Entry) GetName() string {
	return string(e.Name[:])
}
