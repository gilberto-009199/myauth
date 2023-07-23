package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"myauth/application/util"

	"github.com/kbinani/screenshot"
	"github.com/tuotoo/qrcode"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) CaptureScreen() string {

	var img = util.ReadScreen()

	var buf bytes.Buffer

	png.Encode(&buf, img)

	payload := base64.StdEncoding.EncodeToString(buf.Bytes())

	return payload
}

func (a *App) CaptureScreenQRCode(point1 []int, point2 []int) string {

	img, err := screenshot.CaptureRect(image.Rect(point1[0], point1[1], point2[0], point2[1]))
	if err != nil {
		return `{"status":false}`
	}

	buff := new(bytes.Buffer)

	// encode image to buffer
	err = png.Encode(buff, img)
	if err != nil {
		fmt.Println("failed to create buffer", err)
		return `{"status":false}`
	}

	// convert buffer to reader
	reader := bytes.NewReader(buff.Bytes())

	qrmatrix, err := qrcode.Decode(reader)

	if err != nil {
		fmt.Println(err.Error())
		return `{"status":false}`
	}

	message, err := util.ReadOTP(qrmatrix.Content)
	if err != nil {
		fmt.Println(err.Error())
		return `{"status":false}`
	}

	return fmt.Sprintf(`{"status":true,"message":%s}`, message)
}
