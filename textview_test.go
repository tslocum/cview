package cview

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gdamore/tcell"
)

const (
	// 512 bytes
	randomDataSize = 512

	// Write randomData 64 times (32768 bytes) before appending
	appendSetupWriteCount = 64
)

var (
	randomData        = generateRandomData()
	textViewTestCases = generateTestCases()
)

type textViewTestCase struct {
	app      bool
	color    bool
	region   bool
	scroll   bool
	wrap     bool
	wordwrap bool
}

func (c *textViewTestCase) String() string {
	return fmt.Sprintf("Append=%c/Color=%c/Region=%c/Scroll=%c/Wrap=%c/WordWrap=%c", cl(c.app), cl(c.color), cl(c.region), cl(c.scroll), cl(c.wrap), cl(c.wordwrap))
}

func TestTextViewWrite(t *testing.T) {
	t.Parallel()

	for _, c := range textViewTestCases {
		c := c // Capture

		t.Run(c.String(), func(t *testing.T) {
			t.Parallel()

			var (
				tv, _, err   = testTextView(c)
				expectedData []byte
				n            int
			)
			if err != nil {
				t.Error(err)
			}

			if c.app {
				expectedData, err = prepareAppendTextView(tv)
				if err != nil {
					t.Error(err)
				}

				expectedData = append(expectedData, randomData...)
			} else {
				expectedData = randomData
			}

			n, err = tv.Write(randomData)
			if err != nil {
				t.Errorf("failed to write (successfully wrote %d) bytes: %s", n, err)
			} else if n != randomDataSize {
				t.Errorf("failed to write: expected to write %d bytes, wrote %d", randomDataSize, n)
			}

			contents := tv.GetText(false)
			if len(contents) > 0 {
				contents = contents[0 : len(contents)-1] // Remove extra newline
			}
			if len(contents) != len(expectedData) {
				t.Errorf("failed to write: incorrect contents: expected %d bytes, got %d", len(contents), len(expectedData))
			} else if !bytes.Equal([]byte(contents), expectedData) {
				t.Errorf("failed to write: incorrect contents: values do not match")
			}

			tv.Clear()
		})
	}
}

func BenchmarkTextViewWrite(b *testing.B) {
	for _, c := range textViewTestCases {
		c := c // Capture

		b.Run(c.String(), func(b *testing.B) {
			var (
				tv, _, err = testTextView(c)
				n          int
			)
			if err != nil {
				b.Error(err)
			}

			if c.app {
				_, err = prepareAppendTextView(tv)
				if err != nil {
					b.Error(err)
				}
			}

			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				n, err = tv.Write(randomData)
				if err != nil {
					b.Errorf("failed to write (successfully wrote %d) bytes: %s", n, err)
				} else if n != randomDataSize {
					b.Errorf("failed to write: expected to write %d bytes, wrote %d", randomDataSize, n)
				}

				if !c.app {
					b.StopTimer()
					tv.Clear()
					b.StartTimer()
				}
			}
		})
	}
}

func TestTextViewDraw(t *testing.T) {
	t.Parallel()

	for _, c := range textViewTestCases {
		c := c // Capture

		t.Run(c.String(), func(t *testing.T) {
			t.Parallel()

			var (
				tv, sc, err = testTextView(c)
				n           int
			)
			if err != nil {
				t.Error(err)
			}

			if c.app {
				_, err = prepareAppendTextView(tv)
				if err != nil {
					t.Error(err)
				}

				tv.Draw(sc)
			}

			n, err = tv.Write(randomData)
			if err != nil {
				t.Errorf("failed to write (successfully wrote %d) bytes: %s", n, err)
			} else if n != randomDataSize {
				t.Errorf("failed to write: expected to write %d bytes, wrote %d", randomDataSize, n)
			}

			tv.Draw(sc)
		})
	}
}

func BenchmarkTextViewDraw(b *testing.B) {
	for _, c := range textViewTestCases {
		c := c // Capture

		b.Run(c.String(), func(b *testing.B) {
			var (
				tv, sc, err = testTextView(c)
				n           int
			)
			if err != nil {
				b.Error(err)
			}

			if c.app {
				_, err = prepareAppendTextView(tv)
				if err != nil {
					b.Error(err)
				}

				tv.Draw(sc)
			}

			n, err = tv.Write(randomData)
			if err != nil {
				b.Errorf("failed to write (successfully wrote %d) bytes: %s", n, err)
			} else if n != randomDataSize {
				b.Errorf("failed to write: expected to write %d bytes, wrote %d", randomDataSize, n)
			}

			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				tv.Draw(sc)
			}
		})
	}
}

func generateTestCases() []*textViewTestCase {
	var cases []*textViewTestCase
	for i := 0; i < 2; i++ {
		app := i == 1
		for i := 0; i < 2; i++ {
			color := i == 1
			for i := 0; i < 2; i++ {
				region := i == 1
				for i := 0; i < 2; i++ {
					scroll := i == 1
					for i := 0; i < 2; i++ {
						wrap := i == 1
						for i := 0; i < 2; i++ {
							wordwrap := i == 1
							if !wrap && wordwrap {
								continue // WordWrap requires Wrap
							}
							cases = append(cases, &textViewTestCase{app, color, region, scroll, wrap, wordwrap})
						}
					}
				}
			}
		}
	}
	return cases
}

func generateRandomData() []byte {
	var (
		b bytes.Buffer
		r = 33
	)

	for i := 0; i < randomDataSize; i++ {
		if i%80 == 0 && i <= 160 {
			b.WriteRune('\n')
		} else if i%7 == 0 {
			b.WriteRune(' ')
		} else {
			b.WriteRune(rune(r))
		}

		r++
		if r == 127 {
			r = 33
		}
	}

	return b.Bytes()
}

func tvc(tv *TextView, c *textViewTestCase) *TextView {
	return tv.SetDynamicColors(c.color).SetRegions(c.region).SetScrollable(c.scroll).SetWrap(c.wrap).SetWordWrap(c.wordwrap)
}

func cl(v bool) rune {
	if v {
		return 'T'
	}
	return 'F'
}

func testTextView(c *textViewTestCase) (*TextView, tcell.Screen, error) {
	tv := NewTextView()

	sc := tcell.NewSimulationScreen("UTF-8")
	sc.SetSize(80, 24)

	_ = NewApplication().SetRoot(tv, true).SetScreen(sc)

	return tvc(tv, c), sc, nil
}

func prepareAppendTextView(t *TextView) ([]byte, error) {
	var b []byte
	for i := 0; i < appendSetupWriteCount; i++ {
		b = append(b, randomData...)

		n, err := t.Write(randomData)
		if err != nil {
			return nil, fmt.Errorf("failed to write (successfully wrote %d) bytes: %s", n, err)
		} else if n != randomDataSize {
			return nil, fmt.Errorf("failed to write: expected to write %d bytes, wrote %d", randomDataSize, n)
		}
	}

	return b, nil
}
