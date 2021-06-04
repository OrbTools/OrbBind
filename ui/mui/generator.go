package mui

import (
	"math"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/OrbTools/OrbBind/ui/bind"
	"github.com/OrbTools/OrbCommon/gui"
)

//Generate creates a GUI from definition
func Generate(ui *gui.GUI, window fyne.Window, KBS reflect.Value) (*container.AppTabs, func(reflect.Value), func() reflect.Value) {
	keybind := KBS
	tui := container.NewAppTabs()
	for _, page := range ui.Pages[:] {
		var cont *fyne.Container
		pg := page
		switch page.Type {
		case gui.PGrid:
			{
				cont = container.New(layout.NewGridLayout(int(math.Ceil(math.Sqrt(float64(len(page.Keys)))))))
			}
		default:
			{
				cont = container.New(layout.NewGridLayout(int(math.Ceil(math.Sqrt(float64(len(page.Keys)))))))
			}
		}
		for _, key := range page.Keys[:] {
			ky := key
			btn := widget.NewButton(ky.KeyName, func() {
				bp := bind.NewBindPage(ky.KeyID, window, uint16(keybind.Elem().FieldByName(pg.Hive).Index(ky.KeyID).Uint()))
				cont := bp.Create(ky.KeyName)
				ok := func(ok bool) {
					kb := keybind.Elem()
					if ok {
						field := kb.FieldByName(pg.Hive)
						if field.Kind() == reflect.Array {
							field.Index(ky.KeyID).SetUint(uint64(bp.Bind.Bound))
						}
					}
				}
				dialog.ShowCustomConfirm("Bind", "Save", "Cancel", cont, ok, window)
			})
			cont.Add(btn)
		}
		ti := container.NewTabItem(page.Name, cont)
		tui.Append(ti)
	}
	return tui, func(v reflect.Value) { keybind = v }, func() reflect.Value { return keybind }
}
