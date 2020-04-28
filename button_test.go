package cview

import (
	"testing"

	"github.com/gdamore/tcell"
)

const (
	testButtonLabelA = "Hello, world!"
	testButtonLabelB = "Goodnight, moon!"
)

func TestButton(t *testing.T) {
	t.Parallel()

	b, sc, err := testButton()
	if err != nil {
		t.Error(err)
	}
	if b.GetLabel() != testButtonLabelA {
		t.Errorf("failed to initalize Button: incorrect label: expected %s, got %s", testButtonLabelA, b.GetLabel())
	}

	b.SetLabel(testButtonLabelB)
	if b.GetLabel() != testButtonLabelB {
		t.Errorf("failed to update Button: incorrect label: expected %s, got %s", testButtonLabelB, b.GetLabel())
	}

	b.SetLabel(testButtonLabelA)
	if b.GetLabel() != testButtonLabelA {
		t.Errorf("failed to update Button: incorrect label: expected %s, got %s", testButtonLabelA, b.GetLabel())
	}

	b.Draw(sc)
}

func testButton() (*Button, tcell.Screen, error) {
	b := NewButton(testButtonLabelA)

	sc := tcell.NewSimulationScreen("UTF-8")
	sc.SetSize(80, 24)

	_ = NewApplication().SetRoot(b, true).SetScreen(sc)

	return b, sc, nil
}
