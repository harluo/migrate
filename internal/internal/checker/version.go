package checker

import (
	"github.com/harluo/migrate/internal/internal/internal/constraint"
)

type Version[T constraint.Version] interface {
	Version() T
}
