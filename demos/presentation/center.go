package main

import "git.sr.ht/~tslocum/cview"

// Center returns a new primitive which shows the provided primitive in its
// center, given the provided primitive's size.
func Center(width, height int, p cview.Primitive) cview.Primitive {
	return cview.NewFlex().
		AddItem(cview.NewBox(), 0, 1, false).
		AddItem(cview.NewFlex().
			SetDirection(cview.FlexRow).
			AddItem(cview.NewBox(), 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(cview.NewBox(), 0, 1, false), width, 1, true).
		AddItem(cview.NewBox(), 0, 1, false)
}
