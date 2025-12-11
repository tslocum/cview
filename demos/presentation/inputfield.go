package main

import (
	"codeberg.org/tslocum/cview"
	"github.com/gdamore/tcell/v3"
)

// InputField demonstrates the InputField.
func InputField(nextSlide func()) (title string, info string, content cview.Primitive) {
	input := cview.NewInputField()
	input.SetLabel("Enter a number: ")
	input.SetAcceptanceFunc(cview.InputFieldInteger)
	input.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	return "InputField", "", Code(input, 30, 1, "inputfield")
}
