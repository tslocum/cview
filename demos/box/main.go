// Demo code for the Box primitive.
package main

import (
	"github.com/gdamore/tcell"
	"git.sr.ht/~tslocum/cview"
)

func main() {
	box := cview.NewBox().
		SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetTitle("A [red]c[yellow]o[green]l[darkcyan]o[blue]r[darkmagenta]f[red]u[yellow]l[white] [black:red]c[:yellow]o[:green]l[:darkcyan]o[:blue]r[:darkmagenta]f[:red]u[:yellow]l[white:] [::bu]title")
	if err := cview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
