// Demo code for the CheckBox primitive.
package main

import (
	"codeberg.org/tslocum/cview"
)

func main() {
	app := cview.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	checkbox := cview.NewCheckBox()
	checkbox.SetLabel("Hit Enter to check box: ")

	app.SetRoot(checkbox, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
