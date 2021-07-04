package cview

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// TextArea is a multi-line text input.
type TextArea struct {
	*TextView

	setFocus func(Primitive)

	sync.RWMutex
}

// NewTextArea returns a new TextArea object.
func NewTextArea() *TextArea {
	t := &TextArea{
		TextView: NewTextView(),
	}

	t.TextView.SetShowCursor(true)

	return t
}

// Draw draws this primitive onto the screen.
func (t *TextArea) Draw(screen tcell.Screen) {
	if !t.GetVisible() {
		return
	}

	t.Box.Draw(screen)
	t.TextView.Draw(screen)
}

// InputHandler returns the handler for this primitive.
func (t *TextArea) InputHandler() func(event *tcell.EventKey, setFocus func(p Primitive)) {
	return t.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p Primitive)) {
		if t.setFocus == nil {
			t.setFocus = setFocus
		}

		add := func(r rune) bool {
			t.Write([]byte(string(r)))
			return true
		}

		if event.Key() == tcell.KeyRune {
			add(event.Rune())
			return
		} else if event.Key() == tcell.KeyEnter {
			add('\n')
			return
		} else if event.Key() == tcell.KeyBackspace2 {
			b := t.GetBytes(false)
			t.SetBytes(b[:len(b)-1])
			return
		}

		t.TextView.InputHandler()(event, setFocus)
	})
}

// MouseHandler returns the mouse handler for this primitive.
func (t *TextArea) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
	return t.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
		if t.setFocus == nil {
			t.setFocus = setFocus
		}

		x, y := event.Position()
		if !t.InRect(x, y) {
			return false, nil
		}

		return t.TextView.MouseHandler()(action, event, setFocus)
	})
}
