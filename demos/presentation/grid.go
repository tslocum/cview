package main

import (
	"github.com/gdamore/tcell/v2"
	"gitlab.com/tslocum/cview"
)

// Grid demonstrates the grid layout.
func Grid(nextSlide func()) (title string, content cview.Primitive) {
	modalShown := false
	pages := cview.NewPages()

	newPrimitive := func(text string) cview.Primitive {
		tv := cview.NewTextView()
		tv.SetTextAlign(cview.AlignCenter)
		tv.SetText(text)
		tv.SetDoneFunc(func(key tcell.Key) {
			if modalShown {
				nextSlide()
				modalShown = false
			} else {
				pages.ShowPage("modal")
				modalShown = true
			}
		})
		return tv
	}

	menu := newPrimitive("Menu")
	main := newPrimitive("Main content")
	sideBar := newPrimitive("Side Bar")

	grid := cview.NewGrid()
	grid.SetRows(3, 0, 3)
	grid.SetColumns(0, -4, 0)
	grid.SetBorders(true)
	grid.AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, true)
	grid.AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false)
	grid.AddItem(main, 1, 0, 1, 3, 0, 0, false)
	grid.AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false)
	grid.AddItem(main, 1, 1, 1, 1, 0, 100, false)
	grid.AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)

	modal := cview.NewModal()
	modal.SetText("Resize the window to see how the grid layout adapts")
	modal.AddButtons([]string{"Ok"})
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		pages.HidePage("modal")
	})

	pages.AddPage("grid", grid, true, true)
	pages.AddPage("modal", modal, false, false)

	return "Grid", pages
}
