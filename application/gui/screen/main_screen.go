package view

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func MainView() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello"))

	myWindow.Show()
	myApp.Run()
}
