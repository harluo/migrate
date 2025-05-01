package command

import (
	"context"

	"github.com/goexl/log"
	"github.com/harluo/boot"
	"github.com/harluo/migrate/internal/core/internal/command/internal"
	"github.com/harluo/migrate/internal/core/internal/command/internal/argument"
	"github.com/harluo/migrate/internal/internal/core"
	"github.com/harluo/migrate/internal/internal/db"
	"github.com/harluo/migrate/internal/kernel"
)

type Upgrade struct {
	table      *db.Table
	upgrade    *db.Upgrade
	migrations []kernel.Migration

	id      *argument.Id
	pattern *argument.Pattern
	logger  log.Logger
}

func newUpgrade(get internal.Get) *Upgrade {
	return &Upgrade{
		table:      get.Table,
		upgrade:    get.Upgrade,
		migrations: get.Migrations,

		id:      get.Id,
		pattern: get.Pattern,
		logger:  get.Logger,
	}
}

func (u *Upgrade) Run(ctx context.Context) (err error) {
	if ce := u.table.Create(ctx); nil != ce {
		err = ce
	} else {
		err = u.exec(ctx)
	}

	return
}

func (u *Upgrade) Arguments() []boot.Argument {
	return []boot.Argument{
		u.id,
		u.pattern,
	}
}

func (u *Upgrade) Name() string {
	return "downgrade"
}

func (u *Upgrade) Aliases() []string {
	return []string{
		"u",
		"up",
	}
}

func (u *Upgrade) Usage() string {
	return "升级"
}

func (u *Upgrade) Description() string {
	return "升级数据"
}

func (u *Upgrade) exec(ctx context.Context) (err error) {
	for _, migration := range u.getMigrations() {
		err = u.upgrade.Exec(ctx, migration)
		if nil != err {
			break
		}
	}

	return
}

func (u *Upgrade) getMigrations() (migrations []kernel.Migration) {
	migrations = make([]kernel.Migration, 0, len(u.migrations))
	for _, migration := range u.migrations {
		if 0 != u.id.Value && migration.Id() == u.id.Value {
			migrations = append(migrations, migration)
		} else if 0 == u.id.Value && core.NewTyper(migration).Check(u.pattern.Value) {
			migrations = append(migrations, migration)
		}
	}

	return
}
