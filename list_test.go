package cview

import (
	"testing"
)

const (
	listTextA = "Hello, world!"
	listTextB = "Goodnight, moon!"
	listTextC = "Hello, Dolly!"
)

func TestList(t *testing.T) {
	t.Parallel()

	l := NewList()
	if l.GetItemCount() != 0 {
		t.Errorf("failed to initialize List: expected item count 0, got %d", l.GetItemCount())
	} else if l.GetCurrentItem() != 0 {
		t.Errorf("failed to initialize List: expected current item 0, got %d", l.GetCurrentItem())
	}

	l.AddItem(listTextA, listTextB, 'a', nil)
	if l.GetItemCount() != 1 {
		t.Errorf("failed to update List: expected item count 1, got %d", l.GetItemCount())
	} else if l.GetCurrentItem() != 0 {
		t.Errorf("failed to update List: expected current item 0, got %d", l.GetCurrentItem())
	}

	mainText, secondaryText := l.GetItemText(0)
	if mainText != listTextA {
		t.Errorf("failed to update List: expected main text %s, got %s", listTextA, mainText)
	} else if secondaryText != listTextB {
		t.Errorf("failed to update List: expected secondary text %s, got %s", listTextB, secondaryText)
	}

	l.AddItem(listTextB, listTextC, 'a', nil)
	if l.GetItemCount() != 2 {
		t.Errorf("failed to update List: expected item count 1, got %v", l.GetItemCount())
	} else if l.GetCurrentItem() != 0 {
		t.Errorf("failed to update List: expected current item 0, got %v", l.GetCurrentItem())
	}

	mainText, secondaryText = l.GetItemText(1)
	if mainText != listTextB {
		t.Errorf("failed to update List: expected main text %s, got %s", listTextB, mainText)
	} else if secondaryText != listTextC {
		t.Errorf("failed to update List: expected secondary text %s, got %s", listTextC, secondaryText)
	}

	app, err := newTestApp(l)
	if err != nil {
		t.Errorf("failed to initialize Application: %s", err)
	}

	l.Draw(app.screen)
}
