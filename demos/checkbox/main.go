// Demo code for the CheckBox primitive.
package main

import (
	"gitlab.com/tslocum/cview"
)

func main() {
	app := cview.NewApplication()
	checkbox := cview.NewCheckBox().SetLabel("Hit Enter to check box: ")
	if err := app.SetRoot(checkbox, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
