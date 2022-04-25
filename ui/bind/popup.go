package bind

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/OrbTools/OrbBind/keys"
)

//BindingInfo Genral data to use in a channel
type BindingInfo struct {
	Bindid int
	Bound  uint16
}

//Page Binding UI
type Page struct {
	dev    map[string]fyne.CanvasObject
	Bind   BindingInfo
	window fyne.Window
}

//TypeKey event on key
func (bp *Page) TypeKey(e *fyne.KeyEvent) {
	kp := keys.FyneToKeymap(e)
	bp.Bind.Bound = kp.Evdev
	bp.dev["BL"].(*widget.Label).SetText(kp.Code)
}

func (bp *Page) createGrid() *fyne.Container {
	cont := container.New(layout.NewGridLayoutWithColumns(4))
	cont.Add(widget.NewButton("Clear", func() {
		bp.Bind.Bound = 0x0
		bp.dev["BL"].(*widget.Label).SetText(keys.KeyFromEvdev(bp.Bind.Bound).Code)
	}))
	cont.Add(widget.NewButton("TAB", func() {
		bp.TypeKey(&fyne.KeyEvent{Name: fyne.KeyTab})
	}))
	return cont
}

//Create the binding page popup
func (bp *Page) Create(bid string) fyne.CanvasObject {
	bp.dev = make(map[string]fyne.CanvasObject)
	bp.dev["BL"] = widget.NewLabel(keys.KeyFromEvdev(bp.Bind.Bound).Code)
	pop := container.NewVBox(bp.dev["BL"], bp.createGrid())
	bp.window.Canvas().SetOnTypedKey(bp.TypeKey)
	return pop
}

//NewBindPage Create a new bind popup
func NewBindPage(bid int, w fyne.Window, def uint16) *Page {
	p := new(Page)
	p.window = w
	p.Bind.Bindid = bid
	p.Bind.Bound = def
	return p
}
