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

	"codeberg.org/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

const (
	appInfo      = "Next slide: Ctrl-N  Previous: Ctrl-P  Exit: Ctrl-C  (Navigate with your keyboard and mouse)"
	listInfo     = "Next item: J, Down  Previous item: K, Up  Open context menu: Alt+Enter"
	textViewInfo = "Scroll down: J, Down, PageDown  Scroll up: K, Up, PageUp"
	sliderInfo   = "Decrease: H, J, Left, Down  Increase: K, L, Right, Up"
	formInfo     = "Next field: Tab  Previous field: Shift+Tab  Select: Enter"
	windowInfo   = "Windows may be dragged and resized using the mouse."
)

// Slide is a function which returns the slide's title, any applicable
// information and its main primitive, its. It receives a "nextSlide" function
// which can be called to advance the presentation to the next slide.
type Slide func(nextSlide func()) (title string, info string, content cview.Primitive)

// The application.
var app = cview.NewApplication()

// Starting point for the presentation.
func main() {
	defer app.HandlePanic()

	var debugPort int
	flag.IntVar(&debugPort, "debug", 0, "port to serve debug info")
	flag.Parse()

	if debugPort > 0 {
		go func() {
			log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", debugPort), nil))
		}()
	}

	app.EnableMouse(true)

	// The presentation slides.
	slides := []Slide{
		Cover,
		Introduction,
		Colors,
		TextView1,
		TextView2,
		InputField,
		Slider,
		Form,
		Table,
		TreeView,
		Flex,
		Grid,
		Window,
		End,
	}

	panels := cview.NewTabbedPanels()

	// Create the pages for all slides.
	previousSlide := func() {
		slide, _ := strconv.Atoi(panels.GetCurrentTab())
		slide = (slide - 1 + len(slides)) % len(slides)
		panels.SetCurrentTab(strconv.Itoa(slide))
	}
	nextSlide := func() {
		slide, _ := strconv.Atoi(panels.GetCurrentTab())
		slide = (slide + 1) % len(slides)
		panels.SetCurrentTab(strconv.Itoa(slide))
	}

	cursor := 0
	var slideRegions []int
	for index, slide := range slides {
		slideRegions = append(slideRegions, cursor)

		title, info, primitive := slide(nextSlide)

		h := cview.NewTextView()
		if info != "" {
			h.SetDynamicColors(true)
			h.SetText("  [" + cview.ColorHex(cview.Styles.SecondaryTextColor) + "]Info:[-]  " + info)
		}

		// Create a Flex layout that centers the logo and subtitle.
		f := cview.NewFlex()
		f.SetDirection(cview.FlexRow)
		f.AddItem(h, 1, 1, false)
		f.AddItem(primitive, 0, 1, true)

		panels.AddTab(strconv.Itoa(index), title, f)

		cursor += len(title) + 4
	}
	panels.SetCurrentTab("0")

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
	app.SetRoot(panels, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
