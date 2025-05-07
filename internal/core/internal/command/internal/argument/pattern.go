package argument

type Pattern struct {
	Value string `json:"value,omitempty"`
}

func newPattern() *Pattern {
	return new(Pattern)
}

func (*Pattern) Name() string {
	return "pattern"
}

func (*Pattern) Aliases() []string {
	return []string{
		"p",
		"pat",
	}
}

func (p *Pattern) Target() any {
	return &p.Value
}
