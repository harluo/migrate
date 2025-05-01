package core

import (
	"github.com/drone/envsubst/path"
	"github.com/harluo/migrate/internal/internal/checker"
)

type Typer struct {
	data any
}

func NewTyper(data any) *Typer {
	return &Typer{
		data: data,
	}
}

func (t *Typer) Check(pattern string) (checked bool) {
	checked = true
	if converted, ok := t.data.(checker.Name); ok {
		checked, _ = path.Match(pattern, converted.Name())
	}

	return
}

func (t *Typer) Version() (version uint16) {
	switch converted := t.data.(type) {
	case checker.Version[int]:
		version = uint16(converted.Version())
	case checker.Version[int8]:
		version = uint16(converted.Version())
	case checker.Version[int16]:
		version = uint16(converted.Version())
	case checker.Version[int32]:
		version = uint16(converted.Version())
	case checker.Version[int64]:
		version = uint16(converted.Version())
	case checker.Version[uint]:
		version = uint16(converted.Version())
	case checker.Version[uint8]:
		version = uint16(converted.Version())
	case checker.Version[uint16]:
		version = converted.Version()
	case checker.Version[uint32]:
		version = uint16(converted.Version())
	case checker.Version[uint64]:
		version = uint16(converted.Version())
	default:
		version = 1
	}

	return
}
