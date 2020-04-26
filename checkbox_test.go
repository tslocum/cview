package cview

import (
	"testing"

	"github.com/gdamore/tcell"
)

const (
	testCheckBoxLabelA = "Hello, world!"
	testCheckBoxLabelB = "Goodnight, moon!"
)

func TestCheckBox(t *testing.T) {
	t.Parallel()

	c, sc, err := testCheckBox()
	if err != nil {
		t.Error(err)
	}
	if c.IsChecked() {
		t.Errorf("failed to initalize CheckBox: incorrect initial state: expected unchecked, got checked")
	} else if c.GetLabel() != testCheckBoxLabelA {
		t.Errorf("failed to initalize CheckBox: incorrect label: expected %s, got %s", testCheckBoxLabelA, c.GetLabel())
	}

	c.SetLabel(testCheckBoxLabelB)
	if c.GetLabel() != testCheckBoxLabelB {
		t.Errorf("failed to set CheckBox label: incorrect label: expected %s, got %s", testCheckBoxLabelB, c.GetLabel())
	}

	c.SetChecked(true)
	if !c.IsChecked() {
		t.Errorf("failed to update CheckBox state: incorrect state: expected checked, got unchecked")
	}

	c.SetChecked(false)
	if c.IsChecked() {
		t.Errorf("failed to update CheckBox state: incorrect state: expected unchecked, got checked")
	}

	c.Draw(sc)
}

func testCheckBox() (*CheckBox, tcell.Screen, error) {
	c := NewCheckBox()

	sc := tcell.NewSimulationScreen("UTF-8")
	sc.SetSize(80, 24)

	_ = NewApplication().SetRoot(c, true).SetScreen(sc)

	c.SetLabel(testCheckBoxLabelA)

	return c, sc, nil
}
