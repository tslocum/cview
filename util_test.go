package cview

import (
	"github.com/gdamore/tcell"
)

// newTestApp returns a new application connected to a simulation screen.
func newTestApp(root Primitive) (*Application, error) {
	// Initialize simulation screen
	sc := tcell.NewSimulationScreen("UTF-8")
	sc.SetSize(80, 24)

	// Initialize application
	app := NewApplication().
		SetScreen(sc).
		SetRoot(root, true)
	return app, nil
}
