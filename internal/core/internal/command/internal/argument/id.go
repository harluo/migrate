package argument

type Id struct {
	Value uint64 `json:"value,omitempty"`
}

func newId() *Id {
	return new(Id)
}

func (*Id) Name() string {
	return "id"
}

func (*Id) Aliases() []string {
	return []string{
		"i",
		"identify",
	}
}

func (p *Id) Target() any {
	return &p.Value
}
