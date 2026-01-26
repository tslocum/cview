package cview

import (
	"fmt"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/vt"
)

const screenW, screenH = 80, 24

// newTestApp returns a new application connected to a simulation screen.
func newTestApp(root Primitive) (*Application, error) {
	// Initialize simulation screen.
	mt := vt.NewMockTerm(vt.MockOptSize{X: screenW, Y: screenH})
	sc, err := tcell.NewTerminfoScreenFromTty(mt)
	if err != nil {
		return nil, fmt.Errorf("failed to create mock terminal screen: %s", err)
	}
	err = sc.Init()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize mock terminal screen: %s", err)
	}

	// Initialize application.
	app := NewApplication()
	app.SetScreen(sc)
	app.SetRoot(root, true)
	return app, nil
}
