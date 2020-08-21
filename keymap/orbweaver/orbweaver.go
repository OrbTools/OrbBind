package orbweaver

import (
	"encoding/binary"

	"fyne.io/fyne"
	"github.com/minizbot2012/orbbind/keys"
)

//PKM format for altering the keymap
type PKM struct {
	MIP [20]uint16
	SIP [6]uint16
	COL [3]int
}

//SaveIntoKeymap saves an orb
func SaveIntoKeymap(mapped *PKM, file fyne.URIWriteCloser) {
	buf := make([]byte, 2)
	for i := 0; i < 26; i++ {
		if i < 20 {
			binary.LittleEndian.PutUint16(buf, uint16(keys.GetSCForASCII(mapped.MIP[i])))
		} else {
			binary.LittleEndian.PutUint16(buf, uint16(keys.GetSCForASCII(mapped.SIP[i-20])))
		}
		file.Write(buf)
	}
	arr := []byte{byte(mapped.COL[0]), byte(mapped.COL[1]), byte(mapped.COL[2])}
	file.Write(arr)
	file.Close()
}

//LoadFile loads an orb
func LoadFile(file fyne.URIReadCloser) *PKM {
	mapped := &PKM{}
	for i := 0; i < 26; i++ {
		b := make([]byte, 2)
		file.Read(b)
		Asc := keys.GetASCIIForSC(binary.LittleEndian.Uint16(b))
		if i < 26 {
			if i < 20 {
				mapped.MIP[i] = Asc
			} else {
				mapped.SIP[i-20] = Asc
			}
		}
	}
	b := make([]byte, 3)
	file.Read(b)
	mapped.COL[0] = int(b[0])
	mapped.COL[1] = int(b[1])
	mapped.COL[2] = int(b[2])
	file.Close()
	return mapped
}
