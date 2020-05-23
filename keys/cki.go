package keys

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/driver/desktop"
)

//CKIFyneKeyMap Control Keys interop
func CKIFyneKeyMap(e fyne.KeyName) int {
	fmt.Println(e)
	switch e {
	case fyne.KeyUp:
		return CommonToASCII["KEY_UP"]
	case fyne.KeyDown:
		return CommonToASCII["KEY_DOWN"]
	case fyne.KeyLeft:
		return CommonToASCII["KEY_LEFT"]
	case fyne.KeyRight:
		return CommonToASCII["KEY_RIGHT"]
	case fyne.KeyEnd:
		return CommonToASCII["KEY_END"]
	case fyne.KeyHome:
		return CommonToASCII["KEY_HOME"]
	case fyne.KeyInsert:
		return CommonToASCII["KEY_INSERT"]
	case fyne.KeyDelete:
		return CommonToASCII["KEY_DELETE"]
	case fyne.KeyEnter:
		return CommonToASCII["KEY_ENTER"]
	case fyne.KeyTab:
		return CommonToASCII["KEY_TAB"]
	case fyne.KeyEscape:
		return CommonToASCII["KEY_ESC"]
	case fyne.KeyF1:
		return CommonToASCII["KEY_F1"]
	case fyne.KeyF2:
		return CommonToASCII["KEY_F2"]
	case fyne.KeyF3:
		return CommonToASCII["KEY_F3"]
	case fyne.KeyF4:
		return CommonToASCII["KEY_F4"]
	case fyne.KeyF5:
		return CommonToASCII["KEY_F5"]
	case fyne.KeyF6:
		return CommonToASCII["KEY_F6"]
	case fyne.KeyF7:
		return CommonToASCII["KEY_F7"]
	case fyne.KeyF8:
		return CommonToASCII["KEY_F8"]
	case fyne.KeyF9:
		return CommonToASCII["KEY_F9"]
	case fyne.KeyF10:
		return CommonToASCII["KEY_F10"]
	case fyne.KeyF11:
		return CommonToASCII["KEY_F11"]
	case fyne.KeyF12:
		return CommonToASCII["KEY_F12"]
	case fyne.KeyPageDown:
		return CommonToASCII["KEY_PAGEDOWN"]
	case fyne.KeyPageUp:
		return CommonToASCII["KEY_PAGEUP"]
	case desktop.KeyAltLeft:
		return CommonToASCII["KEY_LEFTALT"]
	case desktop.KeyAltRight:
		return CommonToASCII["KEY_RIGHTALT"]
	case desktop.KeyControlLeft:
		return CommonToASCII["KEY_LEFTCTRL"]
	case desktop.KeyControlRight:
		return CommonToASCII["KEY_RIGHTCTRL"]
	case desktop.KeyShiftLeft:
		return CommonToASCII["KEY_LEFTSHIFT"]
	case desktop.KeyShiftRight:
		return CommonToASCII["KEY_RIGHTSHIFT"]
	default:
		return int(e[0])
	}
}

//CKIDetName Determines CTRL key status and returns ascii or name for it
func CKIDetName(e fyne.KeyName) string {
	i := CKIFyneKeyMap(e)
	if i != 0 {
		return ASCIIToCommon[i]
	}
	return "UNK"
}

//CKIDetControl Determines if this is a control key
func CKIDetControl(e fyne.KeyName) bool {
	i := CKIFyneKeyMap(e)
	if i != 0 {
		return true
	}
	return false
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
