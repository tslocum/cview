// Demo code for the Button primitive.
package main

import "git.sr.ht/~tslocum/cview"

func main() {
	app := cview.NewApplication()
	button := cview.NewButton("Hit Enter to close").SetSelectedFunc(func() {
		app.Stop()
	})
	button.SetBorder(true).SetRect(0, 0, 22, 3)
	if err := app.SetRoot(button, false).Run(); err != nil {
		panic(err)
	}
}
