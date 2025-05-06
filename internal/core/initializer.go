package core

import (
	"context"

	"github.com/harluo/migrate/internal/core/internal/command"
	"github.com/harluo/migrate/internal/internal/db"
)

type Initializer struct {
	upgrade *command.Upgrade
	table   *db.Table
}

func newInitializer(upgrade *command.Upgrade, table *db.Table) *Initializer {
	return &Initializer{
		upgrade: upgrade,
		table:   table,
	}
}

func (i *Initializer) Initialize(ctx context.Context) (err error) {
	if ce := i.table.Create(ctx); nil != ce {
		err = ce
	} else if re := i.upgrade.Run(ctx); nil != re {
		err = re
	}

	return
}
