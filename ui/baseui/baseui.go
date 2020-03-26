package baseui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/Minizbot2012/orbbind/keymap/orbweaver"
)

//BasicPage creates a very basic Page
type BasicPage interface {
	Create() *widget.TabItem
}

//PageWithBindings defines a page with bindings
type PageWithBindings interface {
	BasicPage
	SetBinding(*orbweaver.PKM)
}

//DialogPage is a dialog popup
type DialogPage interface {
	Create() *fyne.CanvasObject
}
