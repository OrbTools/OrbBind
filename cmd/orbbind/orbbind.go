//go:generate boxy

package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/Minizbot2012/orbbind/keymap/orbweaver"
	"github.com/Minizbot2012/orbbind/ui/mainpage"
)

func main() {
	ap := app.NewWithID("com.minizbot2012.orbbind")
	window := ap.NewWindow("Orbweaver Rebinding")
	window.Resize(fyne.NewSize(480, 480))
	omap := &orbweaver.PKM{}
	nmp := mainpage.NewMainPage(window, omap)
	tabs := widget.NewTabContainer(nmp.Create())
	tabs.Resize(window.Content().Size())
	window.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("File", fyne.NewMenuItem("Save", func() {
		dialog.ShowFileSave(func(p string) {
			orbweaver.SaveIntoKeymap(omap, p)
		}, window)
	}), fyne.NewMenuItem("Load", func() {
		dialog.ShowFileOpen(func(p string) {
			omap = orbweaver.LoadFile(p)
			nmp.SetBindings(omap)
		}, window)
	}))))
	window.SetContent(tabs)
	window.ShowAndRun()
}
