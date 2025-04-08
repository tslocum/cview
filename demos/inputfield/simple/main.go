// Demo code for the InputField primitive.
package main

import (
	"codeberg.org/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

func main() {
	app := cview.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	inputField := cview.NewInputField()
	inputField.SetLabel("Enter a number: ")
	inputField.SetPlaceholder("E.g. 1234")
	inputField.SetFieldWidth(10)
	inputField.SetAcceptanceFunc(cview.InputFieldInteger)
	inputField.SetDoneFunc(func(key tcell.Key) {
		app.Stop()
	})

	app.SetRoot(inputField, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
