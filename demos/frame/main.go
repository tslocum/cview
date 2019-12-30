// Demo code for the Frame primitive.
package main

import (
	"github.com/gdamore/tcell"
	"git.sr.ht/~tslocum/cview"
)

func main() {
	app := cview.NewApplication()
	frame := cview.NewFrame(cview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
		SetBorders(2, 2, 2, 2, 4, 4).
		AddText("Header left", true, cview.AlignLeft, tcell.ColorWhite).
		AddText("Header middle", true, cview.AlignCenter, tcell.ColorWhite).
		AddText("Header right", true, cview.AlignRight, tcell.ColorWhite).
		AddText("Header second middle", true, cview.AlignCenter, tcell.ColorRed).
		AddText("Footer middle", false, cview.AlignCenter, tcell.ColorGreen).
		AddText("Footer second middle", false, cview.AlignCenter, tcell.ColorGreen)
	if err := app.SetRoot(frame, true).Run(); err != nil {
		panic(err)
	}
}
