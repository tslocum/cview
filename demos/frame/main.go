// Demo code for the Frame primitive.
package main

import (
	"github.com/gdamore/tcell/v2"
	"gitlab.com/tslocum/cview"
)

func main() {
	app := cview.NewApplication()
	frame := cview.NewFrame(cview.NewBox().SetBackgroundColor(tcell.ColorBlue.TrueColor())).
		SetBorders(2, 2, 2, 2, 4, 4).
		AddText("Header left", true, cview.AlignLeft, tcell.ColorWhite.TrueColor()).
		AddText("Header middle", true, cview.AlignCenter, tcell.ColorWhite.TrueColor()).
		AddText("Header right", true, cview.AlignRight, tcell.ColorWhite.TrueColor()).
		AddText("Header second middle", true, cview.AlignCenter, tcell.ColorRed.TrueColor()).
		AddText("Footer middle", false, cview.AlignCenter, tcell.ColorGreen.TrueColor()).
		AddText("Footer second middle", false, cview.AlignCenter, tcell.ColorGreen.TrueColor())
	if err := app.SetRoot(frame, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
