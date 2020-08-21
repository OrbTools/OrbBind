package keys

import (
	"fyne.io/fyne"
)

var (
	//KCToASCII keycode to ascii table
	KCToASCII map[Code]rune
)

//GetASCIIForSC Returns Ascii for Scancode
func GetASCIIForSC(r uint16) uint16 {
	return uint16(KCToASCII[Code(r)])
}

//GetSCForASCII Returns scancode For Ascii
func GetSCForASCII(r uint16) uint16 {
	return uint16(ASCIIToKC[rune(r)])
}

func init() {
	KCToASCII = make(map[Code]rune)
	for x, y := range ASCIIToKC {
		KCToASCII[y] = x
	}
}

//CKIFyneKeyMap Control Keys interop
func CKIFyneKeyMap(e fyne.KeyName) uint16 {
	if val, ok := KCToASCII[FyneToKC[e]]; ok {
		return uint16(val)
	}
	return 0
}

//CKIKeyNameFromKC Key Name from Keycode
func CKIKeyNameFromKC(r uint16) string {
	return Code(r).String()[4:]
}

//CKIKeyNameFromASCII key name from ascii code
func CKIKeyNameFromASCII(r uint16) string {
	return ASCIIToKC[rune(r)].String()[4:]
}
