package cview

import (
	"math"
	"sync"

	"github.com/gdamore/tcell"
)

// ProgressBar indicates the progress of an operation.
type ProgressBar struct {
	*Box

	// Rune to use when rendering the empty area of the progress bar.
	EmptyRune rune

	// Color of the empty area of the progress bar.
	EmptyColor tcell.Color

	// Rune to use when rendering the filled area of the progress bar.
	FilledRune rune

	// Color of the filled area of the progress bar.
	FilledColor tcell.Color

	// If set to true, instead of filling from left to right, the bar is filled
	// from bottom to top.
	Vertical bool

	max      int
	progress int
	*sync.Mutex
}

// NewProgressBar returns a new progress bar.
func NewProgressBar() *ProgressBar {
	return &ProgressBar{
		Box:         NewBox().SetBackgroundColor(Styles.PrimitiveBackgroundColor),
		EmptyRune:   ' ',
		EmptyColor:  Styles.PrimitiveBackgroundColor,
		FilledRune:  tcell.RuneBlock,
		FilledColor: Styles.PrimaryTextColor,
		max:         100,
		Mutex:       new(sync.Mutex),
	}
}

// SetMax sets the progress required to fill the bar.
func (p *ProgressBar) SetMax(max int) {
	p.Lock()
	defer p.Unlock()

	p.max = max
}

// GetMax returns the progress required to fill the bar.
func (p *ProgressBar) GetMax() int {
	p.Lock()
	defer p.Unlock()

	return p.max
}

// AddProgress adds to the current progress.
func (p *ProgressBar) AddProgress(progress int) {
	p.Lock()
	defer p.Unlock()

	p.progress += progress
}

// SetProgress sets the current progress.
func (p *ProgressBar) SetProgress(progress int) {
	p.Lock()
	defer p.Unlock()

	p.progress = progress
}

// GetProgress gets the current progress.
func (p *ProgressBar) GetProgress() int {
	p.Lock()
	defer p.Unlock()

	return p.progress
}

// Complete returns whether the progress bar has been filled.
func (p *ProgressBar) Complete() bool {
	p.Lock()
	defer p.Unlock()

	return p.progress >= p.max
}

// Draw draws this primitive onto the screen.
func (p *ProgressBar) Draw(screen tcell.Screen) {
	p.Lock()
	defer p.Unlock()

	p.Box.Draw(screen)

	x, y, width, height := p.GetInnerRect()

	barSize := height
	maxLength := width
	if p.Vertical {
		barSize = width
		maxLength = height
	}

	barLength := int(math.RoundToEven(float64(maxLength) * (float64(p.progress) / float64(p.max))))
	if barLength > maxLength {
		barLength = maxLength
	}

	for i := 0; i < barSize; i++ {
		for j := 0; j < barLength; j++ {
			if p.Vertical {
				screen.SetContent(x+i, y+(height-1-j), p.FilledRune, nil, tcell.StyleDefault.Foreground(p.FilledColor).Background(p.backgroundColor))
			} else {
				screen.SetContent(x+j, y+i, p.FilledRune, nil, tcell.StyleDefault.Foreground(p.FilledColor).Background(p.backgroundColor))
			}
		}
		for j := barLength; j < maxLength; j++ {
			if p.Vertical {
				screen.SetContent(x+i, y+(height-1-j), p.EmptyRune, nil, tcell.StyleDefault.Foreground(p.EmptyColor).Background(p.backgroundColor))
			} else {
				screen.SetContent(x+j, y+i, p.EmptyRune, nil, tcell.StyleDefault.Foreground(p.EmptyColor).Background(p.backgroundColor))
			}
		}
	}
}
