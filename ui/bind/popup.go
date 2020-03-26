package bind

import (
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/Minizbot2012/orbbind/keys"
	"github.com/Minizbot2012/orbbind/ui/baseui"
)

//BindingInfo Genral data to use in a channel
type BindingInfo struct {
	Bindid int
	Bound  int
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
	if keys.CKIDetControl(e.Name) {
		kp := keys.CKIDetName(e.Name)
		bp.Bind.Bound = keys.CKIFyneKeyMap(e.Name)
		bp.dev["BL"].(*widget.Label).SetText(kp)
	}
}

//TypeRune event one rune type
func (bp *Page) TypeRune(key rune) {
	kp := int(string(strings.ToUpper(string(key)))[0])
	bp.Bind.Bound = kp
	bp.dev["BL"].(*widget.Label).SetText(keys.CKIName(bp.Bind.Bound))
}

func (bp *Page) createGrid() *fyne.Container {
	cont := fyne.NewContainerWithLayout(layout.NewGridLayoutWithColumns(3))
	cont.AddObject(widget.NewButton("Clear", func() {
		bp.dev["BL"].(*widget.Label).SetText(keys.CKIName(0))
		bp.Bind.Bound = 0x0
	}))
	k1 := widget.NewButton("Tab", func() { bp.TypeKey(&fyne.KeyEvent{Name: fyne.KeyTab}) })
	k2 := widget.NewButton("Left Alt", func() { bp.TypeKey(&fyne.KeyEvent{Name: desktop.KeyAltLeft}) })
	k3 := widget.NewButton("Right Alt", func() { bp.TypeKey(&fyne.KeyEvent{Name: desktop.KeyAltRight}) })
	k4 := widget.NewButton("Left Control", func() { bp.TypeKey(&fyne.KeyEvent{Name: desktop.KeyControlLeft}) })
	k5 := widget.NewButton("Right Contorl", func() { bp.TypeKey(&fyne.KeyEvent{Name: desktop.KeyControlRight}) })
	k6 := widget.NewButton("Left Shift", func() { bp.TypeKey(&fyne.KeyEvent{Name: desktop.KeyShiftLeft}) })
	k7 := widget.NewButton("Right Shift", func() { bp.TypeKey(&fyne.KeyEvent{Name: desktop.KeyShiftRight}) })
	cont.AddObject(k1)
	cont.AddObject(k2)
	cont.AddObject(k3)
	cont.AddObject(k4)
	cont.AddObject(k5)
	cont.AddObject(k6)
	cont.AddObject(k7)
	return cont
}

//Create the binding page popup
func (bp *Page) Create(bid string) fyne.CanvasObject {
	bp.dev = make(map[string]fyne.CanvasObject)
	bp.dev["BL"] = widget.NewLabel(keys.CKIName(bp.Bind.Bound))
	pop := widget.NewVBox(bp.dev["BL"], bp.createGrid())
	bp.window.Canvas().SetOnTypedRune(bp.TypeRune)
	bp.window.Canvas().SetOnTypedKey(bp.TypeKey)
	return pop
}

//NewBindPage Create a new bind popup
func NewBindPage(bid int, w fyne.Window, def int) *Page {
	p := &Page{}
	p.window = w
	p.Bind.Bindid = bid
	p.Bind.Bound = def
	return p
}
