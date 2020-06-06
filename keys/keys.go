package keys

import (
	"encoding/json"

	"fyne.io/fyne"
	"fyne.io/fyne/driver/desktop"
	"github.com/minizbot2012/orbbind/box"
)

//Key Codes Map
var (
	ScanToCommon  map[int]string
	ASCIIToCommon map[int]string

	CommonToScan  map[string]int
	CommonToASCII map[string]int
	FyneToASCII   map[fyne.KeyName]int
	ASCIIToFyne   map[int]fyne.KeyName
)

func genStatics() {
	FyneToASCII[fyne.KeyUp] = CommonToASCII["KEY_UP"]
	FyneToASCII[fyne.KeyDown] = CommonToASCII["KEY_DOWN"]
	FyneToASCII[fyne.KeyRight] = CommonToASCII["KEY_RIGHT"]
	FyneToASCII[fyne.KeyLeft] = CommonToASCII["KEY_LEFT"]
	FyneToASCII[fyne.KeyEnd] = CommonToASCII["KEY_END"]
	FyneToASCII[fyne.KeyHome] = CommonToASCII["KEY_HOME"]
	FyneToASCII[fyne.KeyInsert] = CommonToASCII["KEY_INSERT"]
	FyneToASCII[fyne.KeyDelete] = CommonToASCII["KEY_DELETE"]
	FyneToASCII[fyne.KeyEnter] = CommonToASCII["KEY_ENTER"]
	FyneToASCII[fyne.KeyTab] = CommonToASCII["KEY_TAB"]
	FyneToASCII[fyne.KeyEscape] = CommonToASCII["KEY_ESC"]
	FyneToASCII[fyne.KeyF1] = CommonToASCII["KEY_F1"]
	FyneToASCII[fyne.KeyF2] = CommonToASCII["KEY_F2"]
	FyneToASCII[fyne.KeyF3] = CommonToASCII["KEY_F3"]
	FyneToASCII[fyne.KeyF4] = CommonToASCII["KEY_F4"]
	FyneToASCII[fyne.KeyF5] = CommonToASCII["KEY_F5"]
	FyneToASCII[fyne.KeyF6] = CommonToASCII["KEY_F6"]
	FyneToASCII[fyne.KeyF7] = CommonToASCII["KEY_F7"]
	FyneToASCII[fyne.KeyF8] = CommonToASCII["KEY_F8"]
	FyneToASCII[fyne.KeyF9] = CommonToASCII["KEY_F9"]
	FyneToASCII[fyne.KeyF10] = CommonToASCII["KEY_F10"]
	FyneToASCII[fyne.KeyF11] = CommonToASCII["KEY_F11"]
	FyneToASCII[fyne.KeyF12] = CommonToASCII["KEY_F12"]
	FyneToASCII[fyne.KeyPageDown] = CommonToASCII["KEY_PAGEDOWN"]
	FyneToASCII[fyne.KeyPageUp] = CommonToASCII["KEY_PAGEUP"]
	FyneToASCII[fyne.KeySpace] = CommonToASCII["KEY_SPACE"]
	FyneToASCII[desktop.KeyAltLeft] = CommonToASCII["KEY_LEFTALT"]
	FyneToASCII[desktop.KeyAltRight] = CommonToASCII["KEY_RIGHTALT"]
	FyneToASCII[desktop.KeyControlLeft] = CommonToASCII["KEY_LEFTCTRL"]
	FyneToASCII[desktop.KeyControlRight] = CommonToASCII["KEY_RIGHTCTRL"]
	FyneToASCII[desktop.KeyShiftLeft] = CommonToASCII["KEY_LEFTSHIFT"]
	FyneToASCII[desktop.KeyShiftRight] = CommonToASCII["KEY_RIGHTSHIFT"]
}

func genMap() {
	ASCII := box.Get("ascii.json")
	SC := box.Get("scancodes.json")
	json.Unmarshal(ASCII, &CommonToASCII)
	json.Unmarshal(SC, &CommonToScan)
	for k, v := range CommonToASCII {
		ASCIIToCommon[v] = k
	}
	for k, v := range CommonToScan {
		ScanToCommon[v] = k
	}
	genStatics()
	for k, v := range FyneToASCII {
		ASCIIToFyne[v] = k
	}
}

//GetASCIIForSC Returns Ascii for Scancode
func GetASCIIForSC(r int) int {
	return CommonToASCII[ScanToCommon[r]]
}

//GetSCForASCII Returns scancode For Ascii
func GetSCForASCII(r int) int {
	return CommonToScan[ASCIIToCommon[r]]
}

func init() {
	ScanToCommon = make(map[int]string)
	ASCIIToCommon = make(map[int]string)

	CommonToScan = make(map[string]int)
	CommonToASCII = make(map[string]int)

	FyneToASCII = make(map[fyne.KeyName]int)
	ASCIIToFyne = make(map[int]fyne.KeyName)
	genMap()
}
