// Demo code for the InputField primitive.
package main

import (
	"github.com/gdamore/tcell/v2"
	"gitlab.com/tslocum/cview"
)

func main() {
	app := cview.NewApplication()
	inputField := cview.NewInputField().
		SetLabel("Enter a number: ").
		SetPlaceholder("E.g. 1234").
		SetFieldWidth(10).
		SetAcceptanceFunc(cview.InputFieldInteger).
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})
	if err := app.SetRoot(inputField, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
