package util

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kbinani/screenshot"
)

func ReadScreen() {
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("Display not found")
	}

	a := app.New()
	w := a.NewWindow("Test")

	btn := widget.NewButton("Close", w.Hide)
	center := container.NewCenter(btn)

	// https://github.com/search?q=repo%3Afyne-io%2Ffyne%20%20fullscreen&type=code
	w.SetFullScreen(true)
	w.SetContent(center)
	w.ShowAndRun()

	// rect drawing and image load canva - https://github.com/fyne-io/fyne/issues/2924
	// mause move - https://github.com/fyne-io/fyne/issues/2538
}
