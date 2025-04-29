package internal

import (
	"github.com/harluo/boot"
	"github.com/harluo/di"
)

type Put struct {
	di.Put

	Migrate     boot.Command     `group:"commands"`
	Initializer boot.Initializer `group:"initializers"`
}
