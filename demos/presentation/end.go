package main

import (
	"fmt"

	"codeberg.org/tslocum/cview"
	"github.com/gdamore/tcell/v3"
)

// End shows the final slide.
func End(nextSlide func()) (title string, info string, content cview.Primitive) {
	textView := cview.NewTextView()
	textView.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://codeberg.org/tslocum/cview"
	fmt.Fprint(textView, url)
	return "End", "", Center(len(url), 1, textView)
}
