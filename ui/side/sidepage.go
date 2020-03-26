package sidepage

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

//Page Overweave side button configs
type Page struct {
	baseui.PageWithBindings
	binds  *orbweaver.PKM
	dev    map[string]fyne.CanvasObject
	parent fyne.Window
}

//SetBindings Loads bindings from SIP
func (p *Page) SetBindings(o *orbweaver.PKM) {
	p.binds = o
}

//Create Creates the page
func (p *Page) Create() *widget.TabItem {
	p.dev = make(map[string]fyne.CanvasObject)
	strs := []string{"Upper Button", "Dpad Up", "Dpad Right", "Dpad Down", "Dpad Left", "Lower Button"}
	for j := 0; j < 6; j++ {
		id := j + 1
		p.dev["B"+strconv.Itoa(id)] = widget.NewButton(strs[j], func() {
			popup := bind.NewBindPage(id, p.parent, p.binds.SIP[id-1])
			dialog.ShowCustomConfirm("Binding", "Set", "Cancel", popup.Create(string(id)), func(b bool) {
				if b {
					p.binds.SIP[popup.Bind.Bindid-1] = popup.Bind.Bound
				}
			}, p.parent)
		})
	}
	cont := fyne.NewContainerWithLayout(layout.NewBorderLayout(p.dev["B2"], p.dev["B4"], p.dev["B5"], p.dev["B3"]), p.dev["B2"], p.dev["B4"], p.dev["B5"], p.dev["B3"], widget.NewVBox(p.dev["B1"], p.dev["B6"]))
	return widget.NewTabItem("Side Config", cont)
}

//NewSidePage Creates a new side configuration page
func NewSidePage(parent fyne.Window, pkm *orbweaver.PKM) *Page {
	p := &Page{}
	p.binds = pkm
	p.parent = parent
	return p
}
