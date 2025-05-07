package command

import (
	"context"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/harluo/migrate/internal/core/internal/command/internal"
	"github.com/harluo/migrate/internal/core/internal/command/internal/argument"
	"github.com/harluo/migrate/internal/internal/checker"
	"github.com/harluo/migrate/internal/internal/core"
	"github.com/harluo/migrate/internal/internal/db"
	"github.com/harluo/migrate/internal/kernel"
)

type Downgrade struct {
	internal.Base

	table      *db.Table
	downgrade  *db.Downgrade
	migrations []kernel.Migration

	id      *argument.Id
	pattern *argument.Pattern
	logger  log.Logger
}

func newDown(get internal.Get) *Downgrade {
	return &Downgrade{
		table:      get.Table,
		downgrade:  get.Downgrade,
		migrations: get.Migrations,

		id:      get.Id,
		pattern: get.Pattern,
		logger:  get.Logger,
	}
}

func (d *Downgrade) Name() string {
	return "downgrade"
}

func (d *Downgrade) Aliases() []string {
	return []string{
		"d",
		"down",
	}
}

func (d *Downgrade) Usage() string {
	return "降级数据"
}

func (d *Downgrade) Run(ctx context.Context) (err error) {
	if ce := d.table.Create(ctx); nil != ce {
		err = ce
	} else {
		err = d.exec(ctx)
	}

	return
}

func (d *Downgrade) Description() string {
	return "降级"
}

func (d *Downgrade) exec(ctx context.Context) (err error) {
	if migrations, gme := d.getMigrations(); nil != gme {
		err = gme
	} else {
		for _, migration := range migrations {
			err = d.downgrade.Exec(ctx, migration)
			if nil != err {
				break
			}
		}
	}

	return
}

func (d *Downgrade) getMigrations() (migrations map[uint64]kernel.Migration, err error) {
	migrations = make(map[uint64]kernel.Migration, len(d.migrations))
	for _, migration := range d.migrations {
		if _, ok := migration.(checker.Downgrader); !ok {
			continue
		}

		if cached, exists := migrations[migration.Id()]; exists {
			duplicates := d.Migrations(cached, migration)
			err = exception.New().Message("存在重复的数据迁移脚本").Field(field.New("migrations", duplicates)).Build()
		}
		if nil != err {
			return
		}

		if 0 != d.id.Value && migration.Id() == d.id.Value {
			migrations[migration.Id()] = migration
		} else if 0 == d.id.Value && "" != d.pattern.Value && core.NewTyper(migration).Check(d.pattern.Value) {
			migrations[migration.Id()] = migration
		} else {
			migrations[migration.Id()] = migration
		}
	}

	return
}
