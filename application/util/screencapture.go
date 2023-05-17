package util

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/kbinani/screenshot"
)

func ReadScreen(app fyne.App) {
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("Display not found")
	}

	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	// bounds.Dx(), bounds.Dy()

	// Drawing IMG

	w := app.NewWindow("")

	ctx := container.NewMax(canvas.NewImageFromImage(img))

	// https://github.com/search?q=repo%3Afyne-io%2Ffyne%20%20fullscreen&type=code
	w.SetFullScreen(true)
	//w.Resize( fyne.NewSize( bounds.Dx() , bounds.Dy() ) )
	w.SetContent(ctx)
	w.Show()
	w.SetPadded(false)
	// rect drawing and image load canva - https://github.com/fyne-io/fyne/issues/2924
	// mause move - https://github.com/fyne-io/fyne/issues/2538
}
