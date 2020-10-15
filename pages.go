package cview

type page = panel

// Pages is a wrapper around Panels. It is provided for backwards compatibility.
// Application developers should use Panels instead.
type Pages struct {
	*Panels
}

// NewPages returns a new Panels object.
func NewPages() *Pages {
	return &Pages{NewPanels()}
}

// GetPageCount returns the number of panels currently stored in this object.
func (p *Pages) GetPageCount() int {
	return p.GetPanelCount()
}

// AddPage adds a new panel with the given name and primitive.
func (p *Pages) AddPage(name string, item Primitive, resize, visible bool) {
	p.Add(name, item, resize, visible)
}

// AddAndSwitchToPage calls Add(), then SwitchTo() on that newly added panel.
func (p *Pages) AddAndSwitchToPage(name string, item Primitive, resize bool) {
	p.AddAndSwitchTo(name, item, resize)
}

// RemovePage removes the panel with the given name.
func (p *Pages) RemovePage(name string) {
	p.Remove(name)
}

// HasPage returns true if a panel with the given name exists in this object.
func (p *Pages) HasPage(name string) bool {
	return p.Has(name)
}

// ShowPage sets a panel's visibility to "true".
func (p *Pages) ShowPage(name string) {
	p.Show(name)
}

// HidePage sets a panel's visibility to "false".
func (p *Pages) HidePage(name string) {
	p.Hide(name)
}

// SwitchToPage sets a panel's visibility to "true" and all other panels'
// visibility to "false".
func (p *Pages) SwitchToPage(name string) {
	p.SwitchTo(name)
}

// GetFrontPage returns the front-most visible panel.
func (p *Pages) GetFrontPage() (name string, item Primitive) {
	return p.GetFrontPanel()
}
