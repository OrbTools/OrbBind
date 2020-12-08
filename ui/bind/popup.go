package bind

import (
	"fyne.io/fyne"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/minizbot2012/orbbind/keys"
	"github.com/minizbot2012/orbbind/ui/baseui"
)

//BindingInfo Genral data to use in a channel
type BindingInfo struct {
	Bindid int
	Bound  uint16
}

//Page Binding UI
type Page struct {
	baseui.DialogPage
	dev    map[string]fyne.CanvasObject
	Bind   BindingInfo
	window fyne.Window
}

//TypeKey event on key
func (bp *Page) TypeKey(e *fyne.KeyEvent) {
	bp.Bind.Bound = keys.CKIFyneKeyMap(e.Name)
	kp := keys.CKIKeyNameFromKC(bp.Bind.Bound)
	bp.dev["BL"].(*widget.Label).SetText(kp)
}

func (bp *Page) createGrid() *fyne.Container {
	cont := fyne.NewContainerWithLayout(layout.NewGridLayoutWithColumns(4))
	cont.AddObject(widget.NewButton("Clear", func() {
		bp.Bind.Bound = 0x0
		bp.dev["BL"].(*widget.Label).SetText(keys.CKIKeyNameFromKC(bp.Bind.Bound))
	}))
	k1 := widget.NewButton("Tab", func() { bp.TypeKey(&fyne.KeyEvent{Name: fyne.KeyTab}) })
	k3 := widget.NewButton("Left Alt", func() { bp.TypeKey(&fyne.KeyEvent{Name: desktop.KeyAltLeft}) })
	k5 := widget.NewButton("Left Control", func() { bp.TypeKey(&fyne.KeyEvent{Name: desktop.KeyControlLeft}) })
	k7 := widget.NewButton("Left Shift", func() { bp.TypeKey(&fyne.KeyEvent{Name: desktop.KeyShiftLeft}) })
	//k8 := widget.NewButton("Grave (`)", func() { bp.TypeKey(&fyne.KeyEvent{Name: fyne.KeyBackTick}) })
	cont.AddObject(k1)
	cont.AddObject(k3)
	cont.AddObject(k5)
	cont.AddObject(k7)
	return cont
}

//Create the binding page popup
func (bp *Page) Create(bid string) fyne.CanvasObject {
	bp.dev = make(map[string]fyne.CanvasObject)
	bp.dev["BL"] = widget.NewLabel(keys.CKIKeyNameFromKC(bp.Bind.Bound))
	pop := widget.NewVBox(bp.dev["BL"], bp.createGrid())
	bp.window.Canvas().SetOnTypedKey(bp.TypeKey)
	return pop
}

//NewBindPage Create a new bind popup
func NewBindPage(bid int, w fyne.Window, def uint16) *Page {
	p := &Page{}
	p.window = w
	p.Bind.Bindid = bid
	p.Bind.Bound = def
	return p
}
