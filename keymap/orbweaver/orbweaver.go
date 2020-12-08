package orbweaver

import (
	"encoding/binary"

	"fyne.io/fyne"
)

//PKM format for altering the keymap
type PKM struct {
	MIP [20]uint16
	SIP [6]uint16
	COL [3]byte
}

//SaveIntoKeymap saves an orb
func SaveIntoKeymap(mapped *PKM, file fyne.URIWriteCloser) {
	binary.Write(file, binary.LittleEndian, mapped)
	file.Close()
}

//LoadFile loads an orb
func LoadFile(file fyne.URIReadCloser) *PKM {
	mapped := &PKM{}
	binary.Read(file, binary.LittleEndian, mapped)
	file.Close()
	return mapped
}
