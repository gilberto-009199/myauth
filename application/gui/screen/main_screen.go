package view

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainView(app fyne.App) {

	window := app.NewWindow("Hello")
	//fyne.LoadResourceFromPath("23")
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(nil, func() {

		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(nil, func() {}),
		widget.NewToolbarAction(nil, func() {}),
		widget.NewToolbarAction(nil, func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(nil, func() {
			log.Println("Display help")
		}),
	)

	content := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("Content"))

	window.SetContent(content)

	//	window.SetContent(widget.NewLabel("Hello"))

	window.CenterOnScreen()
	window.Show()

}
