package keys

import (
	"fmt"
	"strings"

	"fyne.io/fyne"
)

//CKIFIsCMPLX Determines whether or not to shorten
func CKIFIsCMPLX(e fyne.KeyName) bool {
	_, ok := FyneToASCII[e]
	return ok
}

//CKIAIsCMPLX Determines whether or not to shorten in revo
func CKIAIsCMPLX(i int) bool {
	_, ok := ASCIIToFyne[i]
	return ok
}

//CKIFyneKeyMap Control Keys interop
func CKIFyneKeyMap(e fyne.KeyName) int {
	fmt.Println(e)
	if val, ok := FyneToASCII[e]; ok {
		return val
	}
	return CommonToASCII["KEY_"+strings.ToUpper(string(e))]
}

//CKICommonName returns common name for ascii
func CKICommonName(r int) string {
	return ASCIIToCommon[r]
}

//CKIASCIIIsPrintable returns true if the ascii is in the printable range
func CKIASCIIIsPrintable(r int) bool {
	return r > 32 && r < 127
}

//CKIName returns the character name, or character if printable for item r
func CKIName(r int) string {
	if r == CommonToASCII["KEY_SPACE"] {
		return ASCIIToCommon[r]
	} else if CKIASCIIIsPrintable(r) {
		return string(r)
	}
	return ASCIIToCommon[r]
}
