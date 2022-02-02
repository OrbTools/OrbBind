package main

import (
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/OrbTools/OrbBind/ui/mui"
	"github.com/OrbTools/OrbCommon/devices"
	"github.com/OrbTools/OrbCommon/devices/structs"
)

func main() {
	ap := app.NewWithID("com.minizbot2012.orbbind")
	window := ap.NewWindow("Generic Rebinding Utilities")
	window.SetMaster()
	omap := new(structs.KeyMap)
	main, setter, getter := container.NewAppTabs(), func(reflect.Value) {}, func() reflect.Value { return reflect.ValueOf(0) }
	window.Resize(fyne.NewSize(640, 500))
	devs := fyne.NewMenu("Devices")
	SetDevice := func(dev string) {
		omap = new(structs.KeyMap)
		omap.Device = dev
		omap.Keymap = make([]uint16, devices.DeviceTypes[dev].NumKeys)
		omap.Color = make([]byte, devices.DeviceTypes[dev].NumColor)
		main, setter, getter = mui.Generate(devices.DeviceTypes[dev], window, reflect.ValueOf(omap))
		window.SetContent(main)
	}
	SetDevice("orbweaver")
	for k := range devices.DeviceTypes {
		devs.Items = append(devs.Items, fyne.NewMenuItem(k, func() {
			SetDevice(k)
		}))
	}
	mainMenu := fyne.NewMainMenu(fyne.NewMenu("File", fyne.NewMenuItem("Save", func() {
		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if writer != nil {
				devices.SaveKeymap(writer, getter().Interface())
			}
		}, window)
	}), fyne.NewMenuItem("Load", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if reader != nil {
				omap := devices.LoadKeymap(reader)
				if omap.Device != "" {
					SetDevice(omap.Device)
				} else {
					SetDevice("orbweaver")
				}
				setter(reflect.ValueOf(omap))
			}
		}, window)
	})), devs)
	window.SetMainMenu(mainMenu)
	window.SetContent(main)
	window.ShowAndRun()
}
