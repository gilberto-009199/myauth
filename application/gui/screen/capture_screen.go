package view

import (
	"myauth/application/util"

	"fyne.io/fyne/v2"
)

func CaptureView(app fyne.App) {
	util.ReadScreen(app)
}
