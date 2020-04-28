package cview

import (
	"testing"

	"github.com/gdamore/tcell"
)

const (
	testBoxTitleA = "Hello, world!"
	testBoxTitleB = "Goodnight, moon!"
)

func TestBox(t *testing.T) {
	t.Parallel()

	b, sc, err := testBox()
	if err != nil {
		t.Error(err)
	}
	if b.GetTitle() != "" {
		t.Errorf("failed to initalize Box: incorrect initial state: expected blank title, got %s", b.GetTitle())
	} else if b.border {
		t.Errorf("failed to initalize Box: incorrect initial state: expected no border, got border")
	}

	b.SetTitle(testBoxTitleA)
	if b.GetTitle() != testBoxTitleA {
		t.Errorf("failed to update Box: incorrect title: expected %s, got %s", testBoxTitleA, b.GetTitle())
	}

	b.SetTitle(testBoxTitleB)
	if b.GetTitle() != testBoxTitleB {
		t.Errorf("failed to update Box: incorrect title: expected %s, got %s", testBoxTitleB, b.GetTitle())
	}

	b.SetBorder(true)
	if !b.border {
		t.Errorf("failed to update Box: incorrect state: expected border, got no border")
	}

	b.SetBorder(false)
	if b.border {
		t.Errorf("failed to update Box: incorrect state: expected no border, got border")
	}

	b.Draw(sc)
}

func testBox() (*Box, tcell.Screen, error) {
	b := NewBox()

	sc := tcell.NewSimulationScreen("UTF-8")
	sc.SetSize(80, 24)

	_ = NewApplication().SetRoot(b, true).SetScreen(sc)

	return b, sc, nil
}
