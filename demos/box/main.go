// Demo code for the Box primitive.
package main

import (
	"codeberg.org/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

func main() {
	app := cview.NewApplication()
	defer app.HandlePanic()

	box := cview.NewBox()
	box.SetBorder(true)
	box.SetBorderAttributes(tcell.AttrBold)
	box.SetTitle("A [red]c[yellow]o[green]l[darkcyan]o[blue]r[darkmagenta]f[red]u[yellow]l[white] [black:red]c[:yellow]o[:green]l[:darkcyan]o[:blue]r[:darkmagenta]f[:red]u[:yellow]l[white:] [::bu]title")

	app.SetRoot(box, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
