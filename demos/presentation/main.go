/*
A presentation of the cview package, implemented with cview.

Navigation

The presentation will advance to the next slide when the primitive demonstrated
in the current slide is left (usually by hitting Enter or Escape). Additionally,
the following shortcuts can be used:

  - Ctrl-N: Jump to next slide
  - Ctrl-P: Jump to previous slide
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"gitlab.com/tslocum/cview"
)

// Slide is a function which returns the slide's main primitive and its title.
// It receives a "nextSlide" function which can be called to advance the
// presentation to the next slide.
type Slide func(nextSlide func()) (title string, content cview.Primitive)

// The application.
var app = cview.NewApplication()

// Starting point for the presentation.
func main() {
	var debugPort int
	flag.IntVar(&debugPort, "debug", 0, "port to serve debug info")
	flag.Parse()

	if debugPort > 0 {
		go func() {
			log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", debugPort), nil))
		}()
	}

	// The presentation slides.
	slides := []Slide{
		Cover,
		Introduction,
		Colors,
		TextView1,
		TextView2,
		InputField,
		Form,
		Table,
		TreeView,
		Flex,
		Grid,
		End,
	}

	pages := cview.NewPages()

	// The bottom row has some info on where we are.
	info := cview.NewTextView()
	info.
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			pages.SwitchToPage(added[0])
		})

	// Create the pages for all slides.
	previousSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide - 1 + len(slides)) % len(slides)
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	nextSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide + 1) % len(slides)
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}

	cursor := 0
	var slideRegions []int
	for index, slide := range slides {
		slideRegions = append(slideRegions, cursor)

		title, primitive := slide(nextSlide)
		pages.AddPage(strconv.Itoa(index), primitive, true, index == 0)
		fmt.Fprintf(info, `["%d"][darkcyan] %s [white][""]|`, index, title)

		cursor += len(title) + 4
	}
	info.Highlight("0")

	// Create the main layout.
	layout := cview.NewFlex().
		SetDirection(cview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)

	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			nextSlide()
		} else if event.Key() == tcell.KeyCtrlP {
			previousSlide()
		}
		return event
	})

	// Start the application.
	if err := app.SetRoot(layout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
