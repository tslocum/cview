package main

import (
	"fmt"

	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

// End shows the final slide.
func End(nextSlide func()) (title string, content cview.Primitive) {
	textView := cview.NewTextView()
	textView.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://code.rocketnine.space/tslocum/cview"
	fmt.Fprint(textView, url)
	return "End", Center(len(url), 1, textView)
}
