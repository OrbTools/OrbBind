package keys

// Common Interface 2, natified go
import (
	"fyne.io/fyne"
	"fyne.io/fyne/driver/desktop"
)

//ASCIIToKC -- ASCII Value to a KBD Scancode
var ASCIIToKC = map[rune]Code{
	'a':  CodeA,
	'b':  CodeB,
	'c':  CodeC,
	'd':  CodeD,
	'e':  CodeE,
	'f':  CodeF,
	'g':  CodeG,
	'h':  CodeH,
	'i':  CodeI,
	'k':  CodeK,
	'l':  CodeL,
	'm':  CodeM,
	'n':  CodeN,
	'o':  CodeO,
	'p':  CodeP,
	'q':  CodeQ,
	'r':  CodeR,
	's':  CodeS,
	't':  CodeT,
	'u':  CodeU,
	'v':  CodeV,
	'w':  CodeW,
	'x':  CodeX,
	'y':  CodeY,
	'z':  CodeZ,
	'0':  Code0,
	'1':  Code1,
	'2':  Code2,
	'3':  Code3,
	'4':  Code4,
	'5':  Code5,
	'6':  Code6,
	'7':  Code7,
	'8':  Code8,
	'9':  Code9,
	' ':  CodeSpace,
	'-':  CodeMinus,
	'=':  CodeEqual,
	'[':  CodeLeftBrack,
	']':  CodeRightBrace,
	'\\': CodeBackslash,
	';':  CodeSemicolon,
	'\'': CodeApostrophe,
	'`':  CodeGrave,
	',':  CodeComma,
	'.':  CodeDot,
	'/':  CodeSlash,
}

//FyneToKC - conversion between fyne to ode
var FyneToKC = map[fyne.KeyName]Code{
	fyne.Key0:               Code0,
	fyne.Key1:               Code1,
	fyne.Key2:               Code2,
	fyne.Key3:               Code3,
	fyne.Key4:               Code4,
	fyne.Key5:               Code5,
	fyne.Key6:               Code6,
	fyne.Key7:               Code7,
	fyne.Key8:               Code8,
	fyne.Key9:               Code9,
	fyne.KeyA:               CodeA,
	fyne.KeyB:               CodeB,
	fyne.KeyC:               CodeC,
	fyne.KeyD:               CodeD,
	fyne.KeyE:               CodeE,
	fyne.KeyF:               CodeF,
	fyne.KeyG:               CodeG,
	fyne.KeyH:               CodeH,
	fyne.KeyI:               CodeI,
	fyne.KeyJ:               CodeJ,
	fyne.KeyK:               CodeK,
	fyne.KeyL:               CodeL,
	fyne.KeyM:               CodeM,
	fyne.KeyN:               CodeN,
	fyne.KeyO:               CodeO,
	fyne.KeyP:               CodeP,
	fyne.KeyQ:               CodeQ,
	fyne.KeyR:               CodeR,
	fyne.KeyS:               CodeS,
	fyne.KeyT:               CodeT,
	fyne.KeyU:               CodeU,
	fyne.KeyV:               CodeV,
	fyne.KeyW:               CodeW,
	fyne.KeyX:               CodeX,
	fyne.KeyY:               CodeY,
	fyne.KeyZ:               CodeZ,
	fyne.KeyUp:              CodeUpArrow,
	fyne.KeyDown:            CodeDownArrow,
	fyne.KeyRight:           CodeRightArrow,
	fyne.KeyLeft:            CodeLeftArrow,
	fyne.KeyEnd:             CodeEnd,
	fyne.KeyHome:            CodeHome,
	fyne.KeyInsert:          CodeInsert,
	fyne.KeyDelete:          CodeDelete,
	fyne.KeyEnter:           CodeEnter,
	fyne.KeyTab:             CodeTab,
	fyne.KeyEscape:          CodeESC,
	fyne.KeyF1:              CodeF1,
	fyne.KeyF2:              CodeF2,
	fyne.KeyF3:              CodeF3,
	fyne.KeyF4:              CodeF4,
	fyne.KeyF5:              CodeF5,
	fyne.KeyF6:              CodeF6,
	fyne.KeyF7:              CodeF7,
	fyne.KeyF8:              CodeF8,
	fyne.KeyF9:              CodeF9,
	fyne.KeyF10:             CodeF10,
	fyne.KeyF11:             CodeF11,
	fyne.KeyF12:             CodeF12,
	fyne.KeyPageDown:        CodePageDown,
	fyne.KeyPageUp:          CodePageUp,
	fyne.KeySpace:           CodeSpace,
	fyne.KeyBackTick:        CodeGrave,
	desktop.KeyAltLeft:      CodeLeftAlt,
	desktop.KeyAltRight:     CodeRightAlt,
	desktop.KeyControlLeft:  CodeLeftCntl,
	desktop.KeyControlRight: CodeRightControl,
	desktop.KeyShiftLeft:    CodeLeftShift,
	desktop.KeyShiftRight:   CodeRightShift,
}