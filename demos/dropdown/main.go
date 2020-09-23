// Demo code for the DropDown primitive.
package main

import "gitlab.com/tslocum/cview"

func main() {
	app := cview.NewApplication()
	dropdown := cview.NewDropDown().
		SetLabel("Select an option (hit Enter): ").
		SetOptions(nil,
			cview.NewDropDownOption("First"),
			cview.NewDropDownOption("Second"),
			cview.NewDropDownOption("Third"),
			cview.NewDropDownOption("Fourth"),
			cview.NewDropDownOption("Fifth"))
	if err := app.SetRoot(dropdown, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
