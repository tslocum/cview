package main

import (
	"time"

	"gitlab.com/tslocum/cview"
)

// ProgressBar demonstrates the ProgressBar.
func ProgressBar(nextSlide func()) (title string, content cview.Primitive) {
	grid := cview.NewGrid().SetColumns(-1, 6, 4, 30, -1).SetRows(-1, 12, 4, 4, -1)

	verticalProgressBar := cview.NewProgressBar()
	verticalProgressBar.SetBorder(true)
	verticalProgressBar.Vertical = true

	horizontalProgressBar := cview.NewProgressBar()
	horizontalProgressBar.SetBorder(true)
	horizontalProgressBar.SetMax(150)

	padding := cview.NewTextView()
	grid.AddItem(padding, 0, 0, 1, 5, 0, 0, false)
	grid.AddItem(padding, 1, 0, 1, 1, 0, 0, false)
	grid.AddItem(verticalProgressBar, 1, 1, 2, 1, 0, 0, false)
	grid.AddItem(padding, 1, 2, 1, 1, 0, 0, false)
	grid.AddItem(padding, 2, 0, 1, 5, 0, 0, false)
	grid.AddItem(horizontalProgressBar, 3, 3, 1, 1, 0, 0, false)
	grid.AddItem(padding, 1, 4, 1, 1, 0, 0, false)
	grid.AddItem(padding, 4, 0, 1, 5, 0, 0, false)

	go func() {
		t := time.NewTicker(100 * time.Millisecond)
		for range t.C {
			if verticalProgressBar.Complete() {
				verticalProgressBar.SetProgress(0)
			} else {
				verticalProgressBar.AddProgress(1)
			}

			if horizontalProgressBar.Complete() {
				horizontalProgressBar.SetProgress(0)
			} else {
				horizontalProgressBar.AddProgress(1)
			}

			// Queue draw
			app.QueueUpdateDraw(func() {})
		}
	}()

	return "ProgressBar", grid
}
