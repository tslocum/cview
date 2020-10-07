// Demo code for the Pages primitive.
package main

import (
	"fmt"

	"gitlab.com/tslocum/cview"
)

const pageCount = 5

func main() {
	app := cview.NewApplication()
	app.EnableMouse(true)

	pages := cview.NewPages()
	for page := 0; page < pageCount; page++ {
		func(page int) {
			modal := cview.NewModal()
			modal.SetText(fmt.Sprintf("This is page %d. Choose where to go next.", page+1))
			modal.AddButtons([]string{"Next", "Quit"})
			modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonIndex == 0 {
					pages.SwitchToPage(fmt.Sprintf("page-%d", (page+1)%pageCount))
				} else {
					app.Stop()
				}
			})

			pages.AddPage(fmt.Sprintf("page-%d", page), modal, false, page == 0)
		}(page)
	}

	app.SetRoot(pages, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
