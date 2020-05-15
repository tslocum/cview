package cview

// Key defines the keyboard shortcuts of an application.
type Key struct {
	Cancel []string
	Select []string

	FirstItem []string
	LastItem  []string

	PreviousItem []string
	NextItem     []string

	PreviousField []string
	NextField     []string

	PreviousPage []string
	NextPage     []string

	ShowContextMenu []string
}

// Keys defines the keyboard shortcuts of an application.
var Keys = Key{
	Cancel: []string{"Escape"},
	Select: []string{"Enter", "Ctrl+J"}, // Ctrl+J = keypad enter

	FirstItem: []string{"Home", "g"},
	LastItem:  []string{"End", "G"},

	PreviousItem: []string{"Up", "k"},
	NextItem:     []string{"Down", "j"},

	PreviousField: []string{"Backtab"},
	NextField:     []string{"Tab"},

	PreviousPage: []string{"PageUp"},
	NextPage:     []string{"PageDown"},

	ShowContextMenu: []string{"Alt+Enter"},
}
