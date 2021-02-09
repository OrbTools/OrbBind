package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/OrbTools/OrbBind/ui/baseui"
	"github.com/OrbTools/OrbBind/ui/mainpage"
	"github.com/OrbTools/OrbBind/ui/sidepage"
	"github.com/OrbTools/OrbCommon/devices/orbweaver"
)

func main() {
	ap := app.NewWithID("com.minizbot2012.orbbind")
	window := ap.NewWindow("Orbweaver Rebinding")
	window.SetMaster()

	omap := new(orbweaver.PKM)
	pages := make(map[string]baseui.PageWithBindings)
	pages["main"] = mainpage.NewMainPage(window, omap)
	pages["side"] = sidepage.NewSidePage(window, omap)
	tabs := widget.NewTabContainer(pages["main"].Create(), pages["side"].Create())
	tabs.Resize(fyne.NewSize(640, 480))
	main := tabs
	window.Resize(fyne.NewSize(640, 500))
	mainMenu := fyne.NewMainMenu(fyne.NewMenu("File", fyne.NewMenuItem("Save", func() {
		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if writer != nil {
				orbweaver.SavePKMKeymap(omap, writer)
			}
		}, window)
	}), fyne.NewMenuItem("Load", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if reader != nil {
				omap = orbweaver.LoadPKMKeymap(reader)
				pages["main"].SetBindings(omap)
				pages["side"].SetBindings(omap)
			}
		}, window)
	})))
	window.SetMainMenu(mainMenu)

	window.SetContent(main)
	window.ShowAndRun()
}
