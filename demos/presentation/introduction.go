package main

import "gitlab.com/tslocum/cview"

// Introduction returns a cview.List with the highlights of the cview package.
func Introduction(nextSlide func()) (title string, content cview.Primitive) {
	list := cview.NewList()

	reset := func() {
		list.
			Clear().
			AddItem("A Go package for terminal based UIs", "with a special focus on rich interactive widgets", '1', nextSlide).
			AddItem("Based on github.com/gdamore/tcell", "Like termbox but better (see tcell docs)", '2', nextSlide).
			AddItem("Designed to be simple", `"Hello world" is 5 lines of code`, '3', nextSlide).
			AddItem("Good for data entry", `For charts, use "termui" - for low-level views, use "gocui" - ...`, '4', nextSlide).
			AddItem("Supports context menus", "Right click on one of these items or press Alt+Enter", '5', nextSlide).
			AddItem("Extensive documentation", "Demo code is available for each widget", '6', nextSlide)

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
