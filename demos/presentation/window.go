package main

import (
	"gitlab.com/tslocum/cview"
)

const loremIpsumText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

// Window returns the window page.
func Window(nextSlide func()) (title string, content cview.Primitive) {
	wm := cview.NewWindowManager()

	list := cview.NewList().
		AddItem(cview.NewListItem("Item #1")).
		AddItem(cview.NewListItem("Item #2")).
		AddItem(cview.NewListItem("Item #3")).
		AddItem(cview.NewListItem("Item #4")).
		AddItem(cview.NewListItem("Item #5")).
		AddItem(cview.NewListItem("Item #6")).
		AddItem(cview.NewListItem("Item #7")).
		ShowSecondaryText(false)

	loremIpsum := cview.NewTextView().SetText(loremIpsumText)

	w1 := cview.NewWindow(list).
		SetPosition(2, 2).
		SetSize(10, 7)

	w2 := cview.NewWindow(loremIpsum).
		SetPosition(7, 4).
		SetSize(12, 12)

	w1.SetTitle("List")
	w2.SetTitle("Lorem Ipsum")

	wm.Add(w1, w2)

	return "Window", wm
}
