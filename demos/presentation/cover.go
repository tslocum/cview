package main

import (
	"fmt"
	"strings"

	"git.sr.ht/~tslocum/cview"
	"github.com/gdamore/tcell"
)

const logo = `
 ======= ===  === === ======== ===  ===  ===
===      ===  === === ===      ===  ===  ===
===      ===  === === ======   ===  ===  ===
===       ======  === ===       =========== 
 =======    ==    === ========   ==== ====

`

const (
	subtitle   = `cview - Rich Widgets for Terminal UIs`
	navigation = `Ctrl-N: Next slide    Ctrl-P: Previous slide    Ctrl-C: Exit`
)

// Cover returns the cover page.
func Cover(nextSlide func()) (title string, content cview.Primitive) {
	// What's the size of the logo?
	lines := strings.Split(logo, "\n")
	logoWidth := 0
	logoHeight := len(lines)
	for _, line := range lines {
		if len(line) > logoWidth {
			logoWidth = len(line)
		}
	}
	logoBox := cview.NewTextView().
		SetTextColor(tcell.ColorGreen).
		SetDoneFunc(func(key tcell.Key) {
			nextSlide()
		})
	fmt.Fprint(logoBox, logo)

	// Create a frame for the subtitle and navigation infos.
	frame := cview.NewFrame(cview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(subtitle, true, cview.AlignCenter, tcell.ColorWhite).
		AddText("", true, cview.AlignCenter, tcell.ColorWhite).
		AddText(navigation, true, cview.AlignCenter, tcell.ColorDarkMagenta)

	// Create a Flex layout that centers the logo and subtitle.
	flex := cview.NewFlex().
		SetDirection(cview.FlexRow).
		AddItem(cview.NewBox(), 0, 7, false).
		AddItem(cview.NewFlex().
			AddItem(cview.NewBox(), 0, 1, false).
			AddItem(logoBox, logoWidth, 1, true).
			AddItem(cview.NewBox(), 0, 1, false), logoHeight, 1, true).
		AddItem(frame, 0, 10, false)

	return "Start", flex
}
