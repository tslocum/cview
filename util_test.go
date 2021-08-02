package cview

import (
	"strings"
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// newTestApp returns a new application connected to a simulation screen.
func newTestApp(root Primitive) (*Application, error) {
	// Initialize simulation screen
	sc := tcell.NewSimulationScreen("UTF-8")
	sc.SetSize(80, 24)

	// Initialize application
	app := NewApplication()
	app.SetScreen(sc)
	app.SetRoot(root, true)

	return app, nil
}

func TestWordWrap(t *testing.T) {
	t.Parallel()

	cases := []struct {
		text          string
		width         int
		expectedLines int
	}{
		{
			"first line\n\nthird line",
			80,
			3,
		},
		{
			"first line\n\nthird line",
			6,
			5,
		},
		{
			"string without any newlines",
			80,
			1,
		},
		{
			"string with one newline, string with one newline, string with one>\n<newline, string with one newline",
			32,
			5,
		},
		{
			"[red:-:-]Do you[-:-:-] want to\n\nquit the application?",
			28,
			4,
		},
		{
			Escape("In [Chinese astronomy], the stars that correspond to Gemini are located in two areas: the [White Tiger of the West] (西方白虎, *Xī Fāng Bái Hǔ*) and the [Vermillion Bird of the South] (南方朱雀, *Nán Fāng Zhū Què*)."),
			100,
			3,
		},
	}
	for i, c := range cases {
		lines := WordWrap(c.text, c.width)
		for j, line := range lines {
			if strings.Contains(line, "\n") {
				t.Errorf("case %d, line %d: wrapped line contains newline", i, j)
			}
			strippedWidth := runewidth.StringWidth(string(StripTags([]byte(line), true, true)))
			if strippedWidth > c.width {
				t.Errorf("case %d, line %d: wrapped line exceeds width %d with length %d", i, j, c.width, runewidth.StringWidth(line))
			}
		}
		if len(lines) != c.expectedLines {
			t.Errorf("case %d: expected %d lines, got %d", i, c.expectedLines, len(lines))
		}
	}
}
