package orbweaver

import (
	"encoding/binary"
	"log"
	"os"

	"github.com/Minizbot2012/orbbind/keys"
)

//PKM format for altering the keymap
type PKM struct {
	MIP [20]int
	SIP [6]int
	COL [3]int
}

//SaveIntoKeymap saves an orb
func SaveIntoKeymap(mapped *PKM, path string) {
	fil, _ := os.OpenFile(path, os.O_RDWR, 0)
	for i := 0; i < 26; i++ {
		if i < 26 {
			buf := make([]byte, 2)
			if i < 20 {
				binary.LittleEndian.PutUint16(buf, uint16(keys.GetSCForASCII(int(byte(mapped.MIP[i])))))
			} else {
				binary.LittleEndian.PutUint16(buf, uint16(keys.GetSCForASCII(int(byte(mapped.SIP[i-20])))))
			}
			fil.Write(buf)
		} else {
			arr := []byte{byte(mapped.COL[0]), byte(mapped.COL[1]), byte(mapped.COL[2])}
			fil.Write(arr)
		}
	}
}

//LoadFile loads an orb
func LoadFile(file string) *PKM {
	mapped := &PKM{}
	inf, _ := os.Open(file)
	for i := 0; i < 26; i++ {
		b := make([]byte, 2)
		inf.Read(b)
		Asc := keys.GetASCIIForSC(int(binary.LittleEndian.Uint16(b)))
		log.Println(Asc)
		if i < 26 {
			if i < 20 {
				mapped.MIP[i] = Asc
			} else {
				mapped.SIP[i-20] = Asc
			}
		}
	}
	b := make([]byte, 3)
	inf.Read(b)
	mapped.COL[0] = int(b[0])
	mapped.COL[1] = int(b[1])
	mapped.COL[2] = int(b[2])
	inf.Close()
	return mapped
}
