package util

import (
	"fmt"
	"net/url"
	"os"

	"github.com/tuotoo/qrcode"
)

func ReadQRCode() {
	fi, err := os.Open("qrcode.png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fi.Close()

	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(qrmatrix.Content)
}

func FileQRCode(filename string) (*url.URL, error) {

	fi, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer fi.Close()

	qrmatrix, err := qrcode.Decode(fi)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	message, err := ReadOTPInURLToJSON(qrmatrix.Content)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return message, nil
}
