# Terminal-based user interface toolkit
[![GoDoc](https://godoc.org/git.sr.ht/~tslocum/cview?status.svg)](https://godoc.org/git.sr.ht/~tslocum/cview)
[![builds.sr.ht status](https://builds.sr.ht/~tslocum/cview.svg)](https://builds.sr.ht/~tslocum/cview)

This package is a [fork](https://man.sr.ht/~tslocum/cview/FORK.md) of [tview](https://github.com/rivo/tview).
It allows the creation of rich terminal-based user interfaces.

[![Screenshot of presentation demo](https://git.sr.ht/~tslocum/cview/blob/master/cview.gif)](https://git.sr.ht/~tslocum/cview/tree/master/demos/presentation)

Try the cview presentation demo: ```ssh rocketnine.space -p 20000```

Available widgets:

- __Input forms__ (including __input/password fields__, __drop-down selections__, __checkboxes__, and __buttons__)
- Navigable multi-color __text views__
- Sophisticated navigable __table views__
- Flexible __tree views__
- Selectable __lists__
- __Grid__, __Flexbox__ and __page layouts__
- Modal __message windows__
- An __application__ wrapper

Widgets may be customized and extended to suit any application.

## Installation

```bash
go get git.sr.ht/~tslocum/cview
```

## Hello World

This basic example creates a box titled "Hello, World!" and displays it in your terminal:

```go
package main

import (
	"git.sr.ht/~tslocum/cview"
)

func main() {
	box := cview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := cview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
```

Examples are available in the "demos" subdirectory.

For a presentation highlighting the features of this package, compile and run
the program in the "demos/presentation" subdirectory.

## Documentation

Package documentation is available on [godoc](https://godoc.org/git.sr.ht/~tslocum/cview).

## Dependencies

This package is based on [github.com/gdamore/tcell](https://github.com/gdamore/tcell)
(and its dependencies) and [github.com/rivo/uniseg](https://github.com/rivo/uniseg).

## Support

[CONTRIBUTING.md](https://man.sr.ht/~tslocum/cview/CONTRIBUTING.md) describes how to share
issues, suggestions and patches (pull requests).

cview has two mailing lists:

- [cview-discuss](https://lists.sr.ht/~tslocum/cview-discuss) for general discussion
- [cview-dev](https://lists.sr.ht/~tslocum/cview-dev) for development discussion
