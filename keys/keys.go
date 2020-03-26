package keys

import (
	"encoding/json"

	"github.com/minizbot2012/orbbind/box"
)

//Key Codes Map
var (
	ScanToCommon  map[int]string
	ASCIIToCommon map[int]string

	CommonToScan  map[string]int
	CommonToASCII map[string]int
)

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
	genMap()
}
