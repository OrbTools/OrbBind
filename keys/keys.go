package keys

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/OrbTools/OrbCommon/hid"
)

func FyneToKeymap(e *fyne.KeyEvent) hid.Key {
	switch e.Name {
	case fyne.Key0:
	case fyne.Key1:
	case fyne.Key2:
	case fyne.Key3:
	case fyne.Key4:
	case fyne.Key5:
	case fyne.Key6:
	case fyne.Key7:
	case fyne.Key8:
	case fyne.Key9:
		return hid.GetMappingFromName("DIGIT" + string(e.Name))
	case fyne.KeyF1:
	case fyne.KeyF2:
	case fyne.KeyF3:
	case fyne.KeyF4:
	case fyne.KeyF5:
	case fyne.KeyF6:
	case fyne.KeyF7:
	case fyne.KeyF8:
	case fyne.KeyF9:
	case fyne.KeyF10:
	case fyne.KeyF11:
	case fyne.KeyF12:
		return hid.GetMappingFromName(string(e.Name))
	case fyne.KeyLeft:
	case fyne.KeyRight:
	case fyne.KeyUp:
	case fyne.KeyDown:
		return hid.GetMappingFromName("ARROW_" + strings.ToUpper(string(e.Name)))
	case fyne.KeyBackspace:
	case fyne.KeyBackslash:
	case fyne.KeyEscape:
		return hid.GetMappingFromName(strings.ToUpper(string(e.Name)))
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
	default:
		return hid.GetMappingFromName("US_" + strings.ToUpper(string(e.Name)))
	}
	return hid.GetMappingFromName("USB_RESERVED")
}

func KeyFromEvdev(b uint16) hid.Key {
	return hid.GetMappingFromLinux(b)
}
