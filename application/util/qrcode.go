package util

import (
	"fmt"
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
