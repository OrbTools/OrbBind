package mainpage

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/Minizbot2012/orbbind/keymap/orbweaver"
	"github.com/Minizbot2012/orbbind/ui/baseui"
	"github.com/Minizbot2012/orbbind/ui/bind"
)

//Page is a basic page
type Page struct {
	baseui.PageWithBindings
	binds  *orbweaver.PKM
	dev    map[string]fyne.CanvasObject
	parent fyne.Window
}

func (mp *Page) createButtons() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			id := (j + i*5 + 1)
			mp.dev["B"+strconv.Itoa(id)] = widget.NewButton(strconv.Itoa((j+i*5)+1), func() {
				popup := bind.NewBindPage(id, mp.parent, mp.binds.MIP[id-1])
				dialog.ShowCustomConfirm("Binding", "Set", "Cancel", popup.Create(string(id)), func(b bool) {
					if b {
						mp.binds.MIP[popup.Bind.Bindid-1] = popup.Bind.Bound
					}
				}, mp.parent)
			})
			mp.dev["V"].(*fyne.Container).AddObject(mp.dev["B"+strconv.Itoa(id)])

		}
	}
}

//SetBindings Sets the binding Map
func (mp *Page) SetBindings(o *orbweaver.PKM) {
	mp.binds = o
}

//Create Creates the main binding page
func (mp *Page) Create() *widget.TabItem {
	mp.dev = make(map[string]fyne.CanvasObject)
	layout := layout.NewGridLayout(5)
	mp.dev["V"] = fyne.NewContainerWithLayout(layout)
	mp.createButtons()
	return widget.NewTabItem("Main Bindings", mp.dev["V"])
}

//NewMainPage Creates a new main page
func NewMainPage(parent fyne.Window, pkm *orbweaver.PKM) *Page {
	mp := &Page{}
	mp.binds = pkm
	mp.parent = parent
	return mp
}
