package main

import (
	"github.com/kbinani/screenshot"
)

func main() {
	// Capture each displays.
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("Display not found")
	}

}
