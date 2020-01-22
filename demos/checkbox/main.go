// Demo code for the Checkbox primitive.
package main

import "gitlab.com/tslocum/cview"

func main() {
	app := cview.NewApplication()
	checkbox := cview.NewCheckbox().SetLabel("Hit Enter to check box: ")
	if err := app.SetRoot(checkbox, true).Run(); err != nil {
		panic(err)
	}
}
