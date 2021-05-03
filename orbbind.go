package main

import (
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"github.com/OrbTools/OrbBind/ui/mui"
	"github.com/OrbTools/OrbCommon/devices/orbweaver"
)

func main() {
	ap := app.NewWithID("com.minizbot2012.orbbind")
	window := ap.NewWindow("Orbweaver Rebinding")
	window.SetMaster()

	omap := new(orbweaver.PKM)
	//pages := make(map[string]baseui.PageWithBindings)
	//pages["main"] = mainpage.NewMainPage(window, omap)
	//pages["side"] = sidepage.NewSidePage(window, omap)
	tabs, setter := mui.Generate(orbweaver.GUI, window, reflect.ValueOf(omap))
	//tabs := widget.NewTabContainer(pages["main"].Create(), pages["side"].Create())
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
				omap := orbweaver.LoadPKMKeymap(reader)
				setter(reflect.ValueOf(omap))
				//				pages["main"].SetBindings(omap)
				//				pages["side"].SetBindings(omap)
			}
		}, window)
	})))
	window.SetMainMenu(mainMenu)

	window.SetContent(main)
	window.ShowAndRun()
}
