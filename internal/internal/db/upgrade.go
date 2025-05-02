package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/goexl/db"
	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/harluo/migrate/internal/internal/config"
	"github.com/harluo/migrate/internal/internal/core"
	"github.com/harluo/migrate/internal/internal/db/internal"
	"github.com/harluo/migrate/internal/internal/model"
	"github.com/harluo/migrate/internal/kernel"
)

type Upgrade struct {
	dt        db.Type
	config    *config.Migrate
	migration internal.Migration
}

func newUpgrade(dt db.Type, config *config.Migrate, migration internal.Migration) *Upgrade {
	return &Upgrade{
		dt:        dt,
		config:    config,
		migration: migration,
	}
}

func (u *Upgrade) Exec(ctx context.Context, migration kernel.Migration) (err error) {
	entity := new(model.Migration)
	entity.Id = migration.Id()
	version := core.NewTyper(migration).Version()
	if exists, ge := u.migration.Get(ctx, entity); nil != ge {
		err = ge
	} else if !exists || version > entity.Version {
		err = u.migration.Tx(u.exec(ctx, migration, version))
	}

	return
}

func (u *Upgrade) exec(ctx context.Context, migration kernel.Migration, version uint16) func(*sql.Tx) error {
	return func(tx *sql.Tx) (err error) {
		if ue := migration.Upgrade(ctx); nil != ue {
			err = ue
		} else if affected, ie := u.insert(ctx, tx, &model.Migration{
			Id:          migration.Id(),
			Version:     version,
			Description: migration.Description(),
			Created:     time.Now(),
		}); nil != ie {
			err = ie
		} else if 0 == affected {
			err = exception.New().Message("未插入任何数据").Build()
		}

		return
	}
}

func (u *Upgrade) insert(ctx context.Context, tx *sql.Tx, migration *model.Migration) (affected int64, err error) {
	if query, ise := u.insertSQL(); nil != ise {
		err = ise
	} else if result, ece := tx.ExecContext(
		ctx, query,
		migration.Id, migration.Version, migration.Description, migration.Created,
	); nil != ece {
		err = ece
	} else if nil != result {
		affected, err = result.RowsAffected()
	}

	return
}

func (u *Upgrade) insertSQL() (sql string, err error) {
	switch u.dt {
	case db.TypeMySQL:
		sql = fmt.Sprintf(`INSERT INTO %s (id, version, description) VALUES (?, ?, ?)`, u.config.Table)
	case db.TypePostgres:
		sql = fmt.Sprintf(`INSERT INTO %s (id, version, description) VALUES ($1, $2, $3)`, u.config.Table)
	case db.TypeSQLite:
		sql = fmt.Sprintf(`INSERT INTO %s (id, version, description) VALUES (?, ?, ?)`, u.config.Table)
	case db.TypeOracle:
		sql = fmt.Sprintf(`INSERT INTO %s (id, version, description) VALUES ($1, $2, $3)`, u.config.Table)
	default:
		err = exception.New().Message("不被支持的数据库类型").Field(field.New("type", u.dt)).Build()
	}

	return
}
