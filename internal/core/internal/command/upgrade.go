package command

import (
	"context"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/harluo/boot"
	"github.com/harluo/migrate/internal/core/internal/command/internal"
	"github.com/harluo/migrate/internal/core/internal/command/internal/argument"
	"github.com/harluo/migrate/internal/internal/core"
	"github.com/harluo/migrate/internal/internal/db"
	"github.com/harluo/migrate/internal/kernel"
)

type Upgrade struct {
	internal.Base

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
	return "upgrade"
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
	if migrations, gme := u.getMigrations(); nil != gme {
		err = gme
	} else {
		for _, migration := range migrations {
			err = u.upgrade.Exec(ctx, migration)
			if nil != err {
				break
			}
		}
	}

	return
}

func (u *Upgrade) getMigrations() (migrations map[uint64]kernel.Migration, err error) {
	migrations = make(map[uint64]kernel.Migration, len(u.migrations))
	for _, migration := range u.migrations {
		if cached, exists := migrations[migration.Id()]; exists {
			duplicates := u.Migrations(cached, migration)
			err = exception.New().Message("存在重复的数据迁移脚本").Field(field.New("migrations", duplicates)).Build()
		}
		if nil != err {
			return
		}

		if 0 != u.id.Value && migration.Id() == u.id.Value {
			migrations[migration.Id()] = migration
		} else if 0 == u.id.Value && "" != u.pattern.Value && core.NewTyper(migration).Check(u.pattern.Value) {
			migrations[migration.Id()] = migration
		} else {
			migrations[migration.Id()] = migration
		}
	}

	return
}
