// Demo code for the DropDown primitive.
package main

import "gitlab.com/tslocum/cview"

func main() {
	app := cview.NewApplication()
	dropdown := cview.NewDropDown().
		SetLabel("Select an option (hit Enter): ").
		SetOptions([]string{"First", "Second", "Third", "Fourth", "Fifth"}, nil)
	if err := app.SetRoot(dropdown, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
