package baseui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/OrbTools/OrbCommon/devices/orbweaver"
)

//BasicPage creates a very basic Page
type BasicPage interface {
	Create() *widget.TabItem
}

//PageWithBindings defines a page with bindings
type PageWithBindings interface {
	BasicPage
	SetBindings(*orbweaver.PKM)
}

//DialogPage is a dialog popup
type DialogPage interface {
	Create() *fyne.CanvasObject
}
