package cview

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/gdamore/tcell/v2"
)

// TabbedPanels is a tabbed container for other primitives. The tab switcher
// may be displayed at the top or bottom of the container.
type TabbedPanels struct {
	*Flex
	Switcher *TextView
	panels   *Panels

	tabLabels  map[string]string
	currentTab string

	tabTextColor              tcell.Color
	tabTextColorFocused       tcell.Color
	tabBackgroundColor        tcell.Color
	tabBackgroundColorFocused tcell.Color

	tabDividerStart string
	tabDividerMid   string
	tabDividerEnd   string

	bottomTabSwitcher bool

	width, lastWidth int

	setFocus func(Primitive)

	sync.RWMutex
}

// NewTabbedPanels returns a new TabbedPanels object.
func NewTabbedPanels() *TabbedPanels {
	t := &TabbedPanels{
		Flex:                      NewFlex(),
		Switcher:                  NewTextView(),
		panels:                    NewPanels(),
		tabTextColor:              Styles.PrimaryTextColor,
		tabTextColorFocused:       Styles.InverseTextColor,
		tabBackgroundColor:        ColorUnset,
		tabBackgroundColorFocused: Styles.PrimaryTextColor,
		tabDividerMid:             string(BoxDrawingsDoubleVertical),
		tabDividerEnd:             string(BoxDrawingsLightVertical),
		tabLabels:                 make(map[string]string),
	}

	s := t.Switcher
	s.SetDynamicColors(true)
	s.SetRegions(true)
	s.SetWrap(true)
	s.SetWordWrap(true)
	s.SetHighlightedFunc(func(added, removed, remaining []string) {
		t.SetCurrentTab(added[0])
		if t.setFocus != nil {
			t.setFocus(t.panels)
		}
	})

	f := t.Flex
	f.SetDirection(FlexRow)
	f.AddItem(t.Switcher, 1, 1, false)
	f.AddItem(t.panels, 0, 1, true)

	return t
}

// AddTab adds a new tab. Tab names should consist only of letters, numbers
// and spaces.
func (t *TabbedPanels) AddTab(name, label string, item Primitive) {
	t.Lock()
	t.tabLabels[name] = label
	t.Unlock()

	t.panels.AddPanel(name, item, true, false)

	t.updateAll()
}

// RemoveTab removes a tab.
func (t *TabbedPanels) RemoveTab(name, label string, item Primitive) {
	t.panels.RemovePanel(name)

	t.updateAll()
}

// SetCurrentTab sets the currently visible tab.
func (t *TabbedPanels) SetCurrentTab(name string) {
	t.Lock()

	if t.currentTab == name {
		t.Unlock()
		return
	}

	t.currentTab = name

	t.updateAll()

	t.Unlock()

	t.Switcher.Highlight(t.currentTab)
}

// GetCurrentTab returns the currently visible tab.
func (t *TabbedPanels) GetCurrentTab() string {
	t.RLock()
	defer t.RUnlock()
	return t.currentTab
}

// SetTabLabel sets the label of a tab.
func (t *TabbedPanels) SetTabLabel(name, label string) {
	t.Lock()
	defer t.Unlock()

	if t.tabLabels[name] == label {
		return
	}

	t.tabLabels[name] = label
	t.updateTabLabels()
}

// SetTabTextColor sets the color of the tab text.
func (t *TabbedPanels) SetTabTextColor(color tcell.Color) {
	t.Lock()
	defer t.Unlock()
	t.tabTextColor = color
}

// SetTabTextColorFocused sets the color of the tab text when the tab is in focus.
func (t *TabbedPanels) SetTabTextColorFocused(color tcell.Color) {
	t.Lock()
	defer t.Unlock()
	t.tabTextColorFocused = color
}

// SetTabBackgroundColor sets the background color of the tab.
func (t *TabbedPanels) SetTabBackgroundColor(color tcell.Color) {
	t.Lock()
	defer t.Unlock()
	t.tabBackgroundColor = color
}

// SetTabBackgroundColorFocused sets the background color of the tab when the
// tab is in focus.
func (t *TabbedPanels) SetTabBackgroundColorFocused(color tcell.Color) {
	t.Lock()
	defer t.Unlock()
	t.tabBackgroundColorFocused = color
}

// SetTabSwitcherDivider sets the tab switcher divider text. Color tags are supported.
func (t *TabbedPanels) SetTabSwitcherDivider(start, mid, end string) {
	t.Lock()
	defer t.Unlock()
	t.tabDividerStart, t.tabDividerMid, t.tabDividerEnd = start, mid, end
}

// SetTabSwitcherPosition sets the position of the tab switcher.
func (t *TabbedPanels) SetTabSwitcherPosition(bottom bool) {
	t.Lock()
	defer t.Unlock()

	if t.bottomTabSwitcher == bottom {
		return
	}

	t.bottomTabSwitcher = bottom

	f := t.Flex
	f.RemoveItem(t.panels)
	f.RemoveItem(t.Switcher)
	if t.bottomTabSwitcher {
		f.AddItem(t.panels, 0, 1, true)
		f.AddItem(t.Switcher, 1, 1, false)
	} else {
		f.AddItem(t.Switcher, 1, 1, false)
		f.AddItem(t.panels, 0, 1, true)
	}

	t.updateTabLabels()
}

func (t *TabbedPanels) updateTabLabels() {
	if len(t.panels.panels) == 0 {
		t.Switcher.SetText("")
		t.Flex.ResizeItem(t.Switcher, 0, 1)
		return
	}

	var b bytes.Buffer
	b.WriteString(t.tabDividerStart)
	l := len(t.panels.panels)
	for i, panel := range t.panels.panels {
		textColor := t.tabTextColor
		backgroundColor := t.tabBackgroundColor
		if panel.Name == t.currentTab {
			textColor = t.tabTextColorFocused
			backgroundColor = t.tabBackgroundColorFocused
		}
		b.WriteString(fmt.Sprintf(`["%s"][%s:%s] %s [-:-][""]`, panel.Name, ColorHex(textColor), ColorHex(backgroundColor), t.tabLabels[panel.Name]))

		if i == l-1 {
			b.WriteString(t.tabDividerEnd)
		} else {
			b.WriteString(t.tabDividerMid)
		}
	}
	t.Switcher.SetText(b.String())

	reqLines := len(WordWrap(t.Switcher.GetText(true), t.width))
	if reqLines < 1 {
		reqLines = 1
	}
	t.Flex.ResizeItem(t.Switcher, reqLines, 1)
}

func (t *TabbedPanels) updateVisibleTabs() {
	allPanels := t.panels.panels

	var newTab string

	var foundCurrent bool
	for _, panel := range allPanels {
		if panel.Name == t.currentTab {
			newTab = panel.Name
			foundCurrent = true
			break
		}
	}
	if !foundCurrent {
		for _, panel := range allPanels {
			if panel.Name != "" {
				newTab = panel.Name
				break
			}
		}
	}

	if t.currentTab != newTab {
		t.SetCurrentTab(newTab)
		return
	}

	for _, panel := range allPanels {
		if panel.Name == t.currentTab {
			t.panels.ShowPanel(panel.Name)
		} else {
			t.panels.HidePanel(panel.Name)
		}
	}
}

func (t *TabbedPanels) updateAll() {
	t.updateTabLabels()
	t.updateVisibleTabs()
}

// Draw draws this primitive onto the screen.
func (t *TabbedPanels) Draw(screen tcell.Screen) {
	t.Box.Draw(screen)

	_, _, t.width, _ = t.GetInnerRect()
	if t.width != t.lastWidth {
		t.updateTabLabels()
	}
	t.lastWidth = t.width

	t.Flex.Draw(screen)
}

// InputHandler returns the handler for this primitive.
func (t *TabbedPanels) InputHandler() func(event *tcell.EventKey, setFocus func(p Primitive)) {
	return t.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p Primitive)) {
		if t.setFocus == nil {
			t.setFocus = setFocus
		}
		t.Flex.InputHandler()(event, setFocus)
	})
}

// MouseHandler returns the mouse handler for this primitive.
func (t *TabbedPanels) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
	return t.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
		if t.setFocus == nil {
			t.setFocus = setFocus
		}

		x, y := event.Position()
		if !t.InRect(x, y) {
			return false, nil
		}

		if t.Switcher.InRect(x, y) {
			if t.setFocus != nil {
				defer t.setFocus(t.panels)
			}
			defer t.Switcher.MouseHandler()(action, event, setFocus)
			return true, nil
		}

		return t.Flex.MouseHandler()(action, event, setFocus)
	})
}
