package main

import (
	"github.com/gdamore/tcell"
	"git.sr.ht/~tslocum/cview"
)

// Grid demonstrates the grid layout.
func Grid(nextSlide func()) (title string, content cview.Primitive) {
	modalShown := false
	pages := cview.NewPages()

	newPrimitive := func(text string) cview.Primitive {
		return cview.NewTextView().
			SetTextAlign(cview.AlignCenter).
			SetText(text).
			SetDoneFunc(func(key tcell.Key) {
				if modalShown {
					nextSlide()
					modalShown = false
				} else {
					pages.ShowPage("modal")
					modalShown = true
				}
			})
	}

	menu := newPrimitive("Menu")
	main := newPrimitive("Main content")
	sideBar := newPrimitive("Side Bar")

	grid := cview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(0, -4, 0).
		SetBorders(true).
		AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, true).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false).
		AddItem(main, 1, 0, 1, 3, 0, 0, false).
		AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
		AddItem(main, 1, 1, 1, 1, 0, 100, false).
		AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)

	modal := cview.NewModal().
		SetText("Resize the window to see how the grid layout adapts").
		AddButtons([]string{"Ok"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		pages.HidePage("modal")
	})

	pages.AddPage("grid", grid, true, true).
		AddPage("modal", modal, false, false)

	return "Grid", pages
}
