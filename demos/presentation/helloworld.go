package main

import (
	"github.com/gdamore/tcell"
	"git.sr.ht/~tslocum/cview"
)

const helloWorld = `[green]package[white] main

[green]import[white] (
    [red]"git.sr.ht/~tslocum/cview"[white]
)

[green]func[white] [yellow]main[white]() {
    box := cview.[yellow]NewBox[white]().
        [yellow]SetBorder[white](true).
        [yellow]SetTitle[white]([red]"Hello, world!"[white])
    cview.[yellow]NewApplication[white]().
        [yellow]SetRoot[white](box, true).
        [yellow]Run[white]()
}`

// HelloWorld shows a simple "Hello world" example.
func HelloWorld(nextSlide func()) (title string, content cview.Primitive) {
	// We use a text view because we want to capture keyboard input.
	textView := cview.NewTextView().SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	textView.SetBorder(true).SetTitle("Hello, world!")
	return "Hello, world", Code(textView, 30, 10, helloWorld)
}
