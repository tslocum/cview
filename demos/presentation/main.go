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
	"fmt"
	"strconv"

	"git.sr.ht/~tslocum/cview"
	"github.com/gdamore/tcell"
)

// Slide is a function which returns the slide's main primitive and its title.
// It receives a "nextSlide" function which can be called to advance the
// presentation to the next slide.
type Slide func(nextSlide func()) (title string, content cview.Primitive)

// The application.
var app = cview.NewApplication()

// Starting point for the presentation.
func main() {
	// The presentation slides.
	slides := []Slide{
		Cover,
		Introduction,
		HelloWorld,
		InputField,
		Form,
		TextView1,
		TextView2,
		Table,
		TreeView,
		Flex,
		Grid,
		Colors,
		End,
	}

	// The bottom row has some info on where we are.
	info := cview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false)

	// Create the pages for all slides.
	currentSlide := 0
	info.Highlight(strconv.Itoa(currentSlide))
	pages := cview.NewPages()
	previousSlide := func() {
		currentSlide = (currentSlide - 1 + len(slides)) % len(slides)
		info.Highlight(strconv.Itoa(currentSlide)).
			ScrollToHighlight()
		pages.SwitchToPage(strconv.Itoa(currentSlide))
	}
	nextSlide := func() {
		currentSlide = (currentSlide + 1) % len(slides)
		info.Highlight(strconv.Itoa(currentSlide)).
			ScrollToHighlight()
		pages.SwitchToPage(strconv.Itoa(currentSlide))
	}

	cursor := 0
	var slideRegions []int
	for index, slide := range slides {
		slideRegions = append(slideRegions, cursor)

		title, primitive := slide(nextSlide)
		pages.AddPage(strconv.Itoa(index), primitive, true, index == currentSlide)
		fmt.Fprintf(info, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, title)

		cursor += len(title) + 4
	}

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

	app.EnableMouse()

	var screenHeight int

	app.SetAfterResizeFunc(func(screen tcell.Screen) {
		_, screenHeight = screen.Size()
	})

	app.SetMouseCapture(func(event *cview.EventMouse) *cview.EventMouse {
		atX, atY := event.Position()
		if event.Action()&cview.MouseClick != 0 && atY == screenHeight-1 {
			slideClicked := -1
			for i, region := range slideRegions {
				if atX >= region {
					slideClicked = i
				}
			}
			if slideClicked >= 0 {
				currentSlide = slideClicked
				info.Highlight(strconv.Itoa(currentSlide)).
					ScrollToHighlight()
				pages.SwitchToPage(strconv.Itoa(currentSlide))
			}
			return nil
		}

		return event
	})

	// Start the application.
	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}
