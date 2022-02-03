package keys

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/OrbTools/OrbCommon/hid"
	"github.com/OrbTools/OrbCommon/hid/structs"
)

func FyneToKeymap(e *fyne.KeyEvent) structs.Key {
	switch e.Name {
	case fyne.Key0, fyne.Key1, fyne.Key2, fyne.Key3, fyne.Key4, fyne.Key5, fyne.Key6, fyne.Key7, fyne.Key8, fyne.Key9:
		return hid.GetMappingFromName("DIGIT" + string(e.Name))
	case fyne.KeyF1, fyne.KeyF2, fyne.KeyF3, fyne.KeyF4, fyne.KeyF5, fyne.KeyF6, fyne.KeyF7, fyne.KeyF8, fyne.KeyF9, fyne.KeyF10, fyne.KeyF11, fyne.KeyF12:
		return hid.GetMappingFromName(string(e.Name))
	case fyne.KeyLeft, fyne.KeyRight, fyne.KeyUp, fyne.KeyDown:
		return hid.GetMappingFromName("ARROW_" + strings.ToUpper(string(e.Name)))
	case fyne.KeyBackspace, fyne.KeyEscape, fyne.KeySpace:
		return hid.GetMappingFromName(strings.ToUpper(string(e.Name)))
	case fyne.KeyBackslash:
		return hid.GetMappingFromName("BACKSLASH")
	case fyne.KeyBackTick:
		return hid.GetMappingFromName("BACKQUOTE")
	case fyne.KeyMinus:
		return hid.GetMappingFromName("MINUS")
	case fyne.KeyEqual:
		return hid.GetMappingFromName("EQUAL")
	case fyne.KeySemicolon:
		return hid.GetMappingFromName("SEMICOLON")
	case fyne.KeyReturn:
		return hid.GetMappingFromName("ENTER")
	case desktop.KeyAltLeft:
		return hid.GetMappingFromName("ALT_LEFT")
	case desktop.KeyAltRight:
		return hid.GetMappingFromName("ALT_RIGHT")
	case desktop.KeyControlLeft:
		return hid.GetMappingFromName("CONTROL_LEFT")
	case desktop.KeyControlRight:
		return hid.GetMappingFromName("CONTROL_RIGHT")
	case desktop.KeyShiftLeft:
		return hid.GetMappingFromName("SHIFT_LEFT")
	case desktop.KeyShiftRight:
		return hid.GetMappingFromName("SHIFT_RIGHT")
	case fyne.KeyTab:
		return hid.GetMappingFromName("TAB")
	case fyne.KeyComma:
		return hid.GetMappingFromName("COMMA")
	default:
		return hid.GetMappingFromName("US_" + strings.ToUpper(string(e.Name)))
	}
	//return hid.GetMappingFromName("USB_RESERVED")
}

func KeyFromEvdev(b uint16) structs.Key {
	return hid.GetMappingFromLinux(b)
}
