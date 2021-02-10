package mui

import (
	"math"
	"reflect"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/OrbTools/OrbBind/ui/bind"
	"github.com/OrbTools/OrbCommon/gui"
)

//Generate creates a GUI from definition
func Generate(ui *gui.GUI, window fyne.Window, KBS reflect.Value) (*widget.TabContainer, func(reflect.Value)) {
	keybind := KBS
	tui := widget.NewTabContainer()
	for _, page := range ui.Pages[:] {
		var cont *fyne.Container
		pg := page
		switch page.Type {
		case gui.PGrid:
			{
				cont = fyne.NewContainer()
				cont.Layout = layout.NewGridLayout(int(math.Ceil(math.Sqrt(float64(len(page.Keys))))))
			}
		default:
			{
				cont = fyne.NewContainer()
				cont.Layout = layout.NewGridLayout(int(math.Ceil(math.Sqrt(float64(len(page.Keys))))))
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
							println(field.Index(ky.KeyID).Uint())
							println(ky.KeyID)
						}
					}
				}
				dialog.ShowCustomConfirm("Bind", "Save", "Cancel", cont, ok, window)
			})
			cont.AddObject(btn)
		}
		ti := widget.NewTabItem(page.Name, cont)
		tui.Append(ti)
	}
	return tui, func(v reflect.Value) { keybind = v }
}
