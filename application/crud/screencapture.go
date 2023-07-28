package crud

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"myauth/application/model"
	"myauth/application/util"

	"github.com/kbinani/screenshot"
	"github.com/tuotoo/qrcode"
)

// Section CAPTURE QRCode in Screen
func (a *CrudToken) CaptureScreen(indexMonitor int) string {

	var img = readScreen(indexMonitor)

	var buf bytes.Buffer

	png.Encode(&buf, img)

	payload := base64.StdEncoding.EncodeToString(buf.Bytes())

	return payload
}

func (a *CrudToken) CaptureScreenQRCode(point1 []int, point2 []int) string {

	img, err := screenshot.CaptureRect(image.Rect(point1[0], point1[1], point2[0], point2[1]))
	if err != nil {
		return model.NewMessage(false, nil).ToJSON()
	}

	buff := new(bytes.Buffer)

	// encode image to buffer
	err = png.Encode(buff, img)
	if err != nil {
		fmt.Println("failed to create buffer", err)
		return model.NewMessage(false, nil).ToJSON()
	}

	// convert buffer to reader
	reader := bytes.NewReader(buff.Bytes())

	qrmatrix, err := qrcode.Decode(reader)

	if err != nil {
		fmt.Println(err.Error())
		return model.NewMessage(false, nil).ToJSON()
	}

	message, err := util.ReadOTPInURLToJSON(qrmatrix.Content)
	if err != nil {
		fmt.Println(err.Error())
		return model.NewMessage(false, nil).ToJSON()
	}

	return model.NewMessage(false, TokenResponse{
		Name:     message.Query().Get("issuer"),
		Algoritm: a.appService.Settings.AlgoritmDefault,
		Secret:   message.Query().Get("secret"),
	}).ToJSON()
}

func readScreen(indexMonitor int) *image.RGBA {
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("Display not found")
	}

	bounds := screenshot.GetDisplayBounds(indexMonitor)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	return img
}
