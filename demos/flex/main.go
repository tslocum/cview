// Demo code for the Flex primitive.
package main

import (
	"gitlab.com/tslocum/cview"
)

func main() {
	app := cview.NewApplication()
	flex := cview.NewFlex().
		AddItem(cview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
		AddItem(cview.NewFlex().SetDirection(cview.FlexRow).
			AddItem(cview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
			AddItem(cview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
			AddItem(cview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
		AddItem(cview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
