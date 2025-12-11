package main

import (
	"fmt"

	"codeberg.org/tslocum/cview"
	"github.com/gdamore/tcell/v3"
)

// Slider demonstrates the Slider.
func Slider(nextSlide func()) (title string, info string, content cview.Primitive) {
	slider := cview.NewSlider()
	slider.SetLabel("Volume:   0%")
	slider.SetChangedFunc(func(value int) {
		slider.SetLabel(fmt.Sprintf("Volume: %3d%%", value))
	})
	slider.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	return "Slider", sliderInfo, Code(slider, 30, 1, "slider")
}
