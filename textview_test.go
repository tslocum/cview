package cview

import (
	"bytes"
	"fmt"
	"testing"
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
				tv           = tvc(c)
				expectedData []byte
				n            int
				err          error
			)

			if c.app {
				expectedData, err = prepareAppendTextView(tv)
				if err != nil {
					t.Errorf("failed to prepare append TextView: %s", err)
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
				tv  = tvc(c)
				n   int
				err error
			)

			if c.app {
				_, err = prepareAppendTextView(tv)
				if err != nil {
					b.Errorf("failed to prepare append TextView: %s", err)
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

func BenchmarkTextViewIndex(b *testing.B) {
	for _, c := range textViewTestCases {
		c := c // Capture

		b.Run(c.String(), func(b *testing.B) {
			var (
				tv  = tvc(c)
				n   int
				err error
			)

			_, err = prepareAppendTextView(tv)
			if err != nil {
				b.Errorf("failed to prepare append TextView: %s", err)
			}

			n, err = tv.Write(randomData)
			if err != nil {
				b.Errorf("failed to write (successfully wrote %d) bytes: %s", n, err)
			} else if n != randomDataSize {
				b.Errorf("failed to write: expected to write %d bytes, wrote %d", randomDataSize, n)
			}

			tv.index = nil
			tv.reindexBuffer(80)

			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				tv.index = nil
				tv.reindexBuffer(80)
			}
		})
	}
}

func TestTextViewGetText(t *testing.T) {
	t.Parallel()

	tv := NewTextView()
	tv.SetDynamicColors(true)
	tv.SetRegions(true)

	n, err := tv.Write(randomData)
	if err != nil {
		t.Errorf("failed to write (successfully wrote %d) bytes: %s", n, err)
	} else if n != randomDataSize {
		t.Errorf("failed to write: expected to write %d bytes, wrote %d", randomDataSize, n)
	}

	suffix := []byte(`["start"]outer[b]inner[-]outer[""]`)
	suffixStripped := []byte("outerinnerouter")

	n, err = tv.Write(suffix)
	if err != nil {
		t.Errorf("failed to write (successfully wrote %d) bytes: %s", n, err)
	}

	if !bytes.Equal(tv.GetBytes(false), append(randomData, suffix...)) {
		t.Error("failed to get un-stripped text: unexpected suffix")
	}

	if !bytes.Equal(tv.GetBytes(true), append(randomData, suffixStripped...)) {
		t.Error("failed to get text stripped text: unexpected suffix")
	}
}

func BenchmarkTextViewGetText(b *testing.B) {
	for _, c := range textViewTestCases {
		c := c // Capture

		if c.app {
			continue // Skip for this benchmark
		}

		b.Run(c.String(), func(b *testing.B) {
			var (
				tv  = tvc(c)
				n   int
				err error
				v   []byte
			)

			_, err = prepareAppendTextView(tv)
			if err != nil {
				b.Errorf("failed to prepare append TextView: %s", err)
			}

			n, err = tv.Write(randomData)
			if err != nil {
				b.Errorf("failed to write (successfully wrote %d) bytes: %s", n, err)
			} else if n != randomDataSize {
				b.Errorf("failed to write: expected to write %d bytes, wrote %d", randomDataSize, n)
			}

			v = tv.GetBytes(true)

			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				v = tv.GetBytes(true)
			}

			_ = v
		})
	}
}

type textViewResult struct {
	x         int
	y         int
	primary   rune
	combining []rune
	width     int
}

type textViewRegionsTestCase struct {
	text    string           // Text to test with.
	normal  []textViewResult // How the text should appear normally.
	escaped []textViewResult // How the text should appear when escaped.
}

var textViewHelloWorldResult = []textViewResult{
	{x: 0, y: 0, primary: 'H', combining: nil, width: 1},
	{x: 1, y: 0, primary: 'e', combining: nil, width: 1},
	{x: 2, y: 0, primary: 'l', combining: nil, width: 1},
	{x: 3, y: 0, primary: 'l', combining: nil, width: 1},
	{x: 4, y: 0, primary: 'o', combining: nil, width: 1},
	{x: 5, y: 0, primary: ',', combining: nil, width: 1},
	{x: 6, y: 0, primary: ' ', combining: nil, width: 1},
	{x: 7, y: 0, primary: 'w', combining: nil, width: 1},
	{x: 8, y: 0, primary: 'o', combining: nil, width: 1},
	{x: 9, y: 0, primary: 'r', combining: nil, width: 1},
	{x: 10, y: 0, primary: 'l', combining: nil, width: 1},
	{x: 11, y: 0, primary: 'd', combining: nil, width: 1},
	{x: 12, y: 0, primary: '!', combining: nil, width: 1},
}

var textViewRegionsTestCases = []textViewRegionsTestCase{
	{
		text:    `Hello, world!`,
		normal:  textViewHelloWorldResult,
		escaped: textViewHelloWorldResult,
	}, {
		text: "[TEST\033[0m]\033[36mTEST",
		normal: []textViewResult{
			{x: 0, y: 0, primary: '[', combining: nil, width: 1},
			{x: 1, y: 0, primary: 'T', combining: nil, width: 1},
			{x: 2, y: 0, primary: 'E', combining: nil, width: 1},
			{x: 3, y: 0, primary: 'S', combining: nil, width: 1},
			{x: 4, y: 0, primary: 'T', combining: nil, width: 1},
			{x: 5, y: 0, primary: ']', combining: nil, width: 1},
			{x: 6, y: 0, primary: 'T', combining: nil, width: 1},
			{x: 7, y: 0, primary: 'E', combining: nil, width: 1},
			{x: 8, y: 0, primary: 'S', combining: nil, width: 1},
			{x: 9, y: 0, primary: 'T', combining: nil, width: 1},
		},
		escaped: []textViewResult{
			{x: 0, y: 0, primary: '[', combining: nil, width: 1},
			{x: 1, y: 0, primary: 'T', combining: nil, width: 1},
			{x: 2, y: 0, primary: 'E', combining: nil, width: 1},
			{x: 3, y: 0, primary: 'S', combining: nil, width: 1},
			{x: 4, y: 0, primary: 'T', combining: nil, width: 1},
			{x: 5, y: 0, primary: '[', combining: nil, width: 1},
			{x: 6, y: 0, primary: ']', combining: nil, width: 1},
			{x: 7, y: 0, primary: 'T', combining: nil, width: 1},
			{x: 8, y: 0, primary: 'E', combining: nil, width: 1},
			{x: 9, y: 0, primary: 'S', combining: nil, width: 1},
			{x: 10, y: 0, primary: 'T', combining: nil, width: 1},
		},
	},
}

func TestTextViewANSI(t *testing.T) {
	t.Parallel()

	const screenW, screenH = 80, 80
	for j := 0; j < 2; j++ {
		for i, c := range textViewRegionsTestCases {
			label := "Normal"
			expectedResult := c.normal
			if j == 1 {
				label = "Escaped"
				expectedResult = c.escaped
			}

			t.Run(fmt.Sprintf("%s/%d", label, i+1), func(t *testing.T) {
				t.Parallel()

				tv := NewTextView()
				tv.SetDynamicColors(true)

				app, err := newTestApp(tv)
				if err != nil {
					t.Errorf("failed to initialize Application: %s", err)
				}
				app.screen.SetSize(screenW, screenH)
				tv.SetRect(0, 0, screenW, screenH)

				content := c.text
				if j == 1 {
					content = Escape(content)
				}

				content = TranslateANSI(content)

				tv.SetText(content)

				tv.Draw(app.screen)
				for y := 0; y < screenH; y++ {
					for x := 0; x < screenW; x++ {
						expected := textViewResult{
							primary: ' ',
							width:   1,
						}
						for _, nc := range expectedResult {
							if nc.x == x && nc.y == y {
								expected = nc
								break
							}
						}

						primary, combining, _, width := app.screen.GetContent(x, y)
						if primary != expected.primary {
							t.Errorf("unexpected primary at %d, %d: expected %c, got %c", x, y, expected.primary, primary)
						}

						var combiningEqual bool
						if len(combining) == len(expected.combining) {
							for i, r := range combining {
								if r != expected.combining[i] {
									break
								}
							}
							combiningEqual = true
						}
						if !combiningEqual {
							t.Errorf("unexpected combining at %d, %d: expected %v, got %v", x, y, expected.combining, combining)
						}
						if width != expected.width {
							t.Errorf("unexpected width at %d, %d: expected %d, got %d", x, y, expected.width, width)
						}
					}
				}
			})
		}
	}
}

func TestTextViewDraw(t *testing.T) {
	t.Parallel()

	for _, c := range textViewTestCases {
		c := c // Capture

		t.Run(c.String(), func(t *testing.T) {
			t.Parallel()

			tv := tvc(c)

			app, err := newTestApp(tv)
			if err != nil {
				t.Errorf("failed to initialize Application: %s", err)
			}

			if c.app {
				_, err = prepareAppendTextView(tv)
				if err != nil {
					t.Errorf("failed to prepare append TextView: %s", err)
				}

				tv.Draw(app.screen)
			}

			n, err := tv.Write(randomData)
			if err != nil {
				t.Errorf("failed to write (successfully wrote %d) bytes: %s", n, err)
			} else if n != randomDataSize {
				t.Errorf("failed to write: expected to write %d bytes, wrote %d", randomDataSize, n)
			}

			tv.Draw(app.screen)
		})
	}
}

func BenchmarkTextViewDraw(b *testing.B) {
	for _, c := range textViewTestCases {
		c := c // Capture

		b.Run(c.String(), func(b *testing.B) {
			tv := tvc(c)

			app, err := newTestApp(tv)
			if err != nil {
				b.Errorf("failed to initialize Application: %s", err)
			}

			if c.app {
				_, err = prepareAppendTextView(tv)
				if err != nil {
					b.Errorf("failed to prepare append TextView: %s", err)
				}

				tv.Draw(app.screen)
			}

			n, err := tv.Write(randomData)
			if err != nil {
				b.Errorf("failed to write (successfully wrote %d) bytes: %s", n, err)
			} else if n != randomDataSize {
				b.Errorf("failed to write: expected to write %d bytes, wrote %d", randomDataSize, n)
			}

			tv.Draw(app.screen)

			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				tv.Draw(app.screen)
			}
		})
	}
}

func TestTextViewMaxLines(t *testing.T) {
	t.Parallel()

	tv := NewTextView()

	// append 100 lines with no limit set:
	for i := 0; i < 100; i++ {
		_, err := tv.Write([]byte(fmt.Sprintf("L%d\n", i)))
		if err != nil {
			t.Errorf("failed to write to TextView: %s", err)
		}
	}

	// retrieve the total text and see we have the 100 lines:
	count := bytes.Count(tv.GetBytes(true), []byte("\n"))
	if count != 100 {
		t.Errorf("expected 100 lines, got %d", count)
	}

	// now set the maximum lines to 20, this should clip the buffer:
	tv.SetMaxLines(20)
	// verify buffer was clipped:
	count = len(bytes.Split(tv.GetBytes(true), []byte("\n")))
	if count != 20 {
		t.Errorf("expected 20 lines, got %d", count)
	}

	// append 100 more lines:
	for i := 100; i < 200; i++ {
		_, err := tv.Write([]byte(fmt.Sprintf("L%d\n", i)))
		if err != nil {
			t.Errorf("failed to write to TextView: %s", err)
		}
	}

	// Sice max lines is set to 20, we should still get 20 lines:
	txt := tv.GetBytes(true)
	lines := bytes.Split(txt, []byte("\n"))
	count = len(lines)
	if count != 20 {
		t.Errorf("expected 20 lines, got %d", count)
	}

	// and those 20 lines should be the last ones:
	if !bytes.Equal(lines[0], []byte("L181")) {
		t.Errorf("expected to get L181, got %s", lines[0])
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

func tvc(c *textViewTestCase) *TextView {
	tv := NewTextView()
	tv.SetDynamicColors(c.color)
	tv.SetRegions(c.region)
	tv.SetScrollable(c.scroll)
	tv.SetWrap(c.wrap)
	tv.SetWordWrap(c.wordwrap)
	return tv
}

func cl(v bool) rune {
	if v {
		return 'Y'
	}
	return 'N'
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
