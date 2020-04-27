package cview

import (
	"testing"

	"github.com/gdamore/tcell"
)

func TestProgressBar(t *testing.T) {
	t.Parallel()

	p, sc, err := testProgressBar()
	if err != nil {
		t.Error(err)
	}
	if p.GetProgress() != 0 {
		t.Errorf("failed to initalize ProgressBar: incorrect initial state: expected 0 progress, got %d", p.GetProgress())
	} else if p.GetMax() != 100 {
		t.Errorf("failed to initalize ProgressBar: incorrect initial state: expected 100 max, got %d", p.GetMax())
	} else if p.Complete() {
		t.Errorf("failed to initalize ProgressBar: incorrect initial state: expected incomplete, got complete")
	}

	p.AddProgress(25)
	if p.GetProgress() != 25 {
		t.Errorf("failed to update ProgressBar: incorrect state: expected 25 progress, got %d", p.GetProgress())
	} else if p.Complete() {
		t.Errorf("failed to update ProgressBar: incorrect state: expected incomplete, got complete")
	}

	p.AddProgress(25)
	if p.GetProgress() != 50 {
		t.Errorf("failed to update ProgressBar: incorrect state: expected 50 progress, got %d", p.GetProgress())
	} else if p.Complete() {
		t.Errorf("failed to update ProgressBar: incorrect state: expected incomplete, got complete")
	}

	p.AddProgress(25)
	if p.GetProgress() != 75 {
		t.Errorf("failed to update ProgressBar: incorrect state: expected 75 progress, got %d", p.GetProgress())
	} else if p.Complete() {
		t.Errorf("failed to update ProgressBar: incorrect state: expected incomplete, got complete")
	}

	p.AddProgress(25)
	if p.GetProgress() != 100 {
		t.Errorf("failed to update ProgressBar: incorrect state: expected 100 progress, got %d", p.GetProgress())
	} else if !p.Complete() {
		t.Errorf("failed to update ProgressBar: incorrect state: expected complete, got incomplete")
	}

	p.SetProgress(0)
	if p.GetProgress() != 0 {
		t.Errorf("failed to update ProgressBar: incorrect state: expected 0 progress, got %d", p.GetProgress())
	} else if p.Complete() {
		t.Errorf("failed to update ProgressBar: incorrect state: expected incomplete, got complete")
	}

	p.Draw(sc)
}

func testProgressBar() (*ProgressBar, tcell.Screen, error) {
	p := NewProgressBar()

	sc := tcell.NewSimulationScreen("UTF-8")
	sc.SetSize(80, 24)

	_ = NewApplication().SetRoot(p, true).SetScreen(sc)

	return p, sc, nil
}
