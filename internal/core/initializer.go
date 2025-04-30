package core

import (
	"context"

	"github.com/harluo/migrate/internal/core/internal/command"
	"github.com/harluo/migrate/internal/internal/db"
)

type Initializer struct {
	up    *command.Up
	table *db.Table
}

func newInitializer(up *command.Up) *Initializer {
	return &Initializer{
		up: up,
	}
}

func (i *Initializer) Initialize(ctx context.Context) (err error) {
	if ce := i.table.Create(ctx); nil != ce {
		err = ce
	} else if re := i.up.Run(ctx); nil != re {
		err = re
	}

	return
}
