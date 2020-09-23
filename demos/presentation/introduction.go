package main

import "gitlab.com/tslocum/cview"

// Introduction returns a cview.List with the highlights of the cview package.
func Introduction(nextSlide func()) (title string, content cview.Primitive) {
	list := cview.NewList()

	reset := func() {
		list.
			Clear().
			AddItem(cview.NewListItem("A Go package for terminal based UIs").SetSecondaryText("with a special focus on rich interactive widgets").SetShortcut('1').SetSelectedFunc(nextSlide)).
			AddItem(cview.NewListItem("Based on github.com/gdamore/tcell").SetSecondaryText("Like termbox but better (see tcell docs)").SetShortcut('2').SetSelectedFunc(nextSlide)).
			AddItem(cview.NewListItem("Designed to be simple").SetSecondaryText(`"Hello world" is 5 lines of code`).SetShortcut('3').SetSelectedFunc(nextSlide)).
			AddItem(cview.NewListItem("Good for data entry").SetSecondaryText(`For charts, use "termui" - for low-level views, use "gocui" - ...`).SetShortcut('4').SetSelectedFunc(nextSlide)).
			AddItem(cview.NewListItem("Supports context menus").SetSecondaryText("Right click on one of these items or press Alt+Enter").SetShortcut('5').SetSelectedFunc(nextSlide)).
			AddItem(cview.NewListItem("Extensive documentation").SetSecondaryText("Demo code is available for each widget").SetShortcut('6').SetSelectedFunc(nextSlide))

		list.ContextMenuList().SetItemEnabled(3, false)
	}

	list.AddContextItem("Delete item", 'i', func(index int) {
		list.RemoveItem(index)

		if list.GetItemCount() == 0 {
			list.ContextMenuList().SetItemEnabled(0, false)
			list.ContextMenuList().SetItemEnabled(1, false)
		}
		list.ContextMenuList().SetItemEnabled(3, true)
	})

	list.AddContextItem("Delete all", 'a', func(index int) {
		list.Clear()

		list.ContextMenuList().SetItemEnabled(0, false)
		list.ContextMenuList().SetItemEnabled(1, false)
		list.ContextMenuList().SetItemEnabled(3, true)
	})

	list.AddContextItem("", 0, nil)

	list.AddContextItem("Reset", 'r', func(index int) {
		reset()

		list.ContextMenuList().SetItemEnabled(0, true)
		list.ContextMenuList().SetItemEnabled(1, true)
		list.ContextMenuList().SetItemEnabled(3, false)
	})

	reset()
	return "Introduction", Center(80, 12, list)
}
