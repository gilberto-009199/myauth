package gui

import (
	view "myauth/application/gui/screen"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func Start() {

	app := app.New()
	app.Settings().SetTheme(theme.LightTheme())
	// app.SetIcon(  )

	view.MainView(app)

	app.Run()
}
