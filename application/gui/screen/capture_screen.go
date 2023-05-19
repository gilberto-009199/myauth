package view

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"

	"fyne.io/fyne/v2"
)

func CaptureView(app fyne.App) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	var command = path.Join(exPath, "libcapture")
	os := runtime.GOOS
	switch os {
	case "windows":
		command += ".exe"
	case "darwin":
		command += ""
	case "linux":
		command = "." + command
	default:
		fmt.Printf("%s.\n", os)
	}

	out, err := exec.Command(command).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Out: %s", out)

}

/*  TENTATIVA COM FYNE PURO(Falha mauseout e mauseclick nao aplicavel simultaneamente)

import (
	"myauth/application/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func CaptureView(app fyne.App) {

	img := util.ReadScreen()

	w := app.NewWindow("")
	imageContainer := canvas.NewImageFromImage(img)

	//var _ desktop.Hoverable = (*desktop.Mouseable)(imageContainer)
	//	Hoverable

	// https://github.com/search?q=repo%3Afyne-io%2Ffyne%20%20fullscreen&type=code

	ctx := container.NewMax(imageContainer)
	//w.Resize( fyne.NewSize( bounds.Dx() , bounds.Dy() ) )

	w.SetContent(ctx)
	// NewRasterWithPixels MauseMove - https://github.com/fyne-io/fyne/issues/418
	// Click tapped	-
	//w.SetFullScreen(true)
	w.SetPadded(false)
	w.Resize(fyne.NewSize(500, 300))
	w.Show()

	// rect drawing and image load canva - https://github.com/fyne-io/fyne/issues/2924
	// mause move - https://github.com/fyne-io/fyne/issues/2538

}

/* FIM DA TENTATIVA COM FYNE */
