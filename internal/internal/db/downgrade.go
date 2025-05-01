package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/goexl/db"
	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/harluo/migrate/internal/internal/config"
	"github.com/harluo/migrate/internal/internal/db/internal"
	"github.com/harluo/migrate/internal/internal/model"
	"github.com/harluo/migrate/internal/kernel"
)

type Downgrade struct {
	dt        db.Type
	config    *config.Migrate
	migration internal.Migration
}

func newDowngrade(dt db.Type, config *config.Migrate, migration internal.Migration) *Downgrade {
	return &Downgrade{
		dt:        dt,
		config:    config,
		migration: migration,
	}
}

func (u *Downgrade) Exec(ctx context.Context, migration kernel.Migration) (err error) {
	entity := new(model.Migration)
	entity.Id = migration.Id()
	if exists, ge := u.migration.Get(ctx, entity); nil != ge {
		err = ge
	} else if exists {
		err = u.migration.Tx(u.exec(ctx, migration))
	}

	return
}

func (u *Downgrade) exec(ctx context.Context, migration kernel.Migration) func(*sql.Tx) error {
	return func(tx *sql.Tx) (err error) {
		if ue := migration.Downgrade(ctx); nil != ue {
			err = ue
		} else if affected, ie := u.delete(ctx, tx, &model.Migration{Id: migration.Id()}); nil != ie {
			err = ie
		} else if 0 == affected {
			err = exception.New().Message("未删除任何数据").Build()
		}

		return
	}
}

func (u *Downgrade) delete(ctx context.Context, tx *sql.Tx, migration *model.Migration) (affected int64, err error) {
	if query, ise := u.deleteSQL(); nil != ise {
		err = ise
	} else if result, ece := tx.ExecContext(ctx, query, migration.Id); nil != ece {
		err = ece
	} else if nil != result {
		affected, err = result.RowsAffected()
	}

	return
}

func (u *Downgrade) deleteSQL() (sql string, err error) {
	switch u.dt {
	case db.TypeMysql:
		sql = fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, u.config.Table)
	case db.TypePostgres:
		sql = fmt.Sprintf(`DELETE FROM %s WHERE id = $1)`, u.config.Table)
	case db.TypeSQLite:
		sql = fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, u.config.Table)
	case db.TypeOracle:
		sql = fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, u.config.Table)
	default:
		err = exception.New().Message("不被支持的数据库类型").Field(field.New("type", u.dt)).Build()
	}

	return
}
