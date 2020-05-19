package cview

import (
	"github.com/gdamore/tcell"
	"gitlab.com/tslocum/cbind"
)

// Key defines the keyboard shortcuts of an application.
type Key struct {
	Select    []string
	SelectAlt []string // SelectAlt is also used when not focusing a text input.
	Cancel    []string

	MoveUp    []string
	MoveDown  []string
	MoveLeft  []string
	MoveRight []string

	MoveFirst         []string
	MoveLast          []string
	MovePreviousField []string
	MoveNextField     []string
	MovePreviousPage  []string
	MoveNextPage      []string

	ShowContextMenu []string
}

// Keys defines the keyboard shortcuts of an application.
var Keys = Key{
	Select:    []string{"Enter", "Ctrl+J"}, // Ctrl+J = keypad enter
	SelectAlt: []string{"Space"},
	Cancel:    []string{"Escape"},

	MoveUp:    []string{"Up", "k"},
	MoveDown:  []string{"Down", "j"},
	MoveLeft:  []string{"Left", "h"},
	MoveRight: []string{"Right", "l"},

	MoveFirst:         []string{"Home", "g"},
	MoveLast:          []string{"End", "G"},
	MovePreviousField: []string{"Backtab"},
	MoveNextField:     []string{"Tab"},
	MovePreviousPage:  []string{"PageUp", "Ctrl+B"},
	MoveNextPage:      []string{"PageDown", "Ctrl+F"},

	ShowContextMenu: []string{"Alt+Enter"},
}

// HitShortcut returns whether the EventKey provided is present in one or more
// sets of keybindings.
func HitShortcut(event *tcell.EventKey, keybindings ...[]string) bool {
	enc, err := cbind.Encode(event.Modifiers(), event.Key(), event.Rune())
	if err != nil {
		return false
	}

	for _, binds := range keybindings {
		for _, key := range binds {
			if key == enc {
				return true
			}
		}
	}

	return false
}
