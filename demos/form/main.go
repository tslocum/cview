// Demo code for the Form primitive.
package main

import (
	"gitlab.com/tslocum/cview"
)

func main() {
	app := cview.NewApplication()
	form := cview.NewForm().
		AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
		AddInputField("First name", "", 20, nil, nil).
		AddInputField("Last name", "", 20, nil, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddCheckBox("", "Age 18+", false, nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(cview.AlignLeft)
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
