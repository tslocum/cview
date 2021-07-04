package main

import (
	"code.rocketnine.space/tslocum/cview"
)

// TextArea demonstrates the TextArea.
func TextArea(nextSlide func()) (title string, info string, content cview.Primitive) {
	t := cview.NewTextArea()
	t.SetBorder(true)
	t.SetTitle("Multi-line text input")
	return "TextArea", "", Center(44, 16, t)
}
