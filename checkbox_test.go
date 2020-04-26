package cview

import (
	"testing"

	"github.com/gdamore/tcell"
)

const (
	testCheckboxLabelA = "Hello, world!"
	testCheckboxLabelB = "Goodnight, moon!"
)

func TestCheckbox(t *testing.T) {
	t.Parallel()

	c, sc, err := testCheckbox()
	if err != nil {
		t.Error(err)
	}
	if c.IsChecked() {
		t.Errorf("failed to initalize checkbox: incorrect initial state: expected unchecked, got checked")
	} else if c.GetLabel() != testCheckboxLabelA {
		t.Errorf("failed to initalize checkbox: incorrect label: expected %s, got %s", testCheckboxLabelA, c.GetLabel())
	}

	c.SetLabel(testCheckboxLabelB)
	if c.GetLabel() != testCheckboxLabelB {
		t.Errorf("failed to set checkbox label: incorrect label: expected %s, got %s", testCheckboxLabelA, c.GetLabel())
	}

	c.SetChecked(true)
	if !c.IsChecked() {
		t.Errorf("failed to update checkbox state: incorrect state: expected checked, got unchecked")
	}

	c.SetChecked(false)
	if c.IsChecked() {
		t.Errorf("failed to update checkbox state: incorrect state: expected unchecked, got checked")
	}

	c.Draw(sc)
}

func testCheckbox() (*Checkbox, tcell.Screen, error) {
	c := NewCheckbox()

	sc := tcell.NewSimulationScreen("UTF-8")
	sc.SetSize(80, 24)

	_ = NewApplication().SetRoot(c, true).SetScreen(sc)

	c.SetLabel(testCheckboxLabelA)

	return c, sc, nil
}
