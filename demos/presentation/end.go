package main

import (
	"fmt"

	"github.com/gdamore/tcell"
	"git.sr.ht/~tslocum/cview"
)

// End shows the final slide.
func End(nextSlide func()) (title string, content cview.Primitive) {
	textView := cview.NewTextView().SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://git.sr.ht/~tslocum/cview"
	fmt.Fprint(textView, url)
	return "End", Center(len(url), 1, textView)
}
