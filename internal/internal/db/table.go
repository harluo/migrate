package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/goexl/db"
	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/harluo/migrate/internal/internal/config"
)

type Table struct {
	dt     db.Type
	db     *sql.DB
	config *config.Migrate
}

func newTable(dt db.Type, db *sql.DB, config *config.Migrate) *Table {
	return &Table{
		dt:     dt,
		db:     db,
		config: config,
	}
}

func (t *Table) Create(ctx context.Context) (err error) {
	if exists, ee := t.exists(); nil != ee {
		err = ee
	} else if !exists {
		err = t.create(ctx)
	}

	return
}

func (t *Table) create(ctx context.Context) (err error) {
	if tx, bte := t.db.BeginTx(ctx, nil); nil != bte {
		err = bte
	} else if cte := t.createWithTx(ctx, tx); nil != cte {
		err = tx.Rollback()
	} else {
		err = tx.Commit()
	}

	return
}

func (t *Table) createWithTx(ctx context.Context, tx *sql.Tx) (err error) {
	if cte := t.createTable(ctx, tx); nil != cte {
		err = cte
	} else if db.TypeMySQL == t.dt {
		err = t.commentMySQL(ctx, tx)
	} else if db.TypePostgres == t.dt {
		err = t.commentPostgresql(ctx, tx)
	} else {
		err = tx.Commit()
	}

	return
}

func (t *Table) commentMySQL(ctx context.Context, tx *sql.Tx) (err error) {
	if _, cte := tx.ExecContext(ctx, fmt.Sprintf("ALTER TABLE %s COMMENT = '数据迁移记录'", t.config.Table)); nil != cte {
		err = cte
	} else if _, cie := tx.ExecContext(ctx, fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN id COMMENT '标识'", t.config.Table)); nil != cie {
		err = cie
	} else if _, cve := tx.ExecContext(ctx, fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN version COMMENT '版本号'", t.config.Table)); nil != cve {
		err = cve
	} else if _, cde := tx.ExecContext(ctx, fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN description COMMENT '本次数据迁移描述信息'", t.config.Table)); nil != cde {
		err = cde
	} else if _, cce := tx.ExecContext(ctx, fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN created COMMENT '本次数据迁移创建时间'", t.config.Table)); nil != cce {
		err = cce
	} else if _, cue := tx.ExecContext(ctx, fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN updated COMMENT '本次数据迁移修改时间'", t.config.Table)); nil != cue {
		err = cue
	}

	return
}

func (t *Table) commentPostgresql(ctx context.Context, tx *sql.Tx) (err error) {
	if _, cte := tx.ExecContext(ctx, fmt.Sprintf("COMMENT ON TABLE %s IS '数据迁移记录'", t.config.Table)); nil != cte {
		err = cte
	} else if _, cie := tx.ExecContext(ctx, fmt.Sprintf("COMMENT ON COLUMN %s.id IS '标识'", t.config.Table)); nil != cie {
		err = cie
	} else if _, cve := tx.ExecContext(ctx, fmt.Sprintf("COMMENT ON COLUMN %s.version IS '版本号'", t.config.Table)); nil != cve {
		err = cve
	} else if _, cde := tx.ExecContext(ctx, fmt.Sprintf("COMMENT ON COLUMN %s.description IS '本次数据迁移描述信息'", t.config.Table)); nil != cde {
		err = cde
	} else if _, cce := tx.ExecContext(ctx, fmt.Sprintf("COMMENT ON COLUMN %s.created IS '本次数据迁移创建时间'", t.config.Table)); nil != cce {
		err = cce
	} else if _, cue := tx.ExecContext(ctx, fmt.Sprintf("COMMENT ON COLUMN %s.updated IS '本次数据迁移修改时间'", t.config.Table)); nil != cue {
		err = cue
	}

	return
}

func (t *Table) createTable(ctx context.Context, tx *sql.Tx) (err error) {
	if smallint, bigint, datetime, varchar, de := t.dialect(); nil != de {
		err = de
	} else {
		_, err = tx.ExecContext(ctx, fmt.Sprintf(`CREATE TABLE %s (
	id %s PRIMARY KEY,

	version %s NOT NULL DEFAULT 1,
	description %s NOT NULL DEFAULT '',

	created %s NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated %s NOT NULL DEFAULT CURRENT_TIMESTAMP
)`, t.config.Table, bigint, smallint, varchar, datetime, datetime))
	}

	return
}

func (t *Table) exists() (exists bool, err error) {
	if query, qe := t.createSQL(); nil != qe {
		err = qe
	} else {
		err = t.db.QueryRow(query, t.config.Table).Scan(&exists)
	}

	return
}

func (t *Table) createSQL() (sql string, err error) {
	switch t.dt {
	case db.TypeMySQL:
		sql = `SELECT COUNT(*) FROM "information_schema"."TABLES" WHERE "TABLE_NAME" = ?`
	case db.TypePostgres:
		sql = `SELECT EXISTS(SELECT FROM "pg_tables" WHERE "schemaname" = 'public' AND "tablename" = $1)`
	case db.TypeSQLite:
		sql = `SELECT "name" FROM "sqlite_master" WHERE "type"='table' AND "name" = ?`
	case db.TypeOracle:
		sql = `SELECT COUNT(*) FROM sys.objects WHERE object_id = OBJECT_ID(?) AND type = 'U'`
	default:
		err = exception.New().Message("不被支持的数据库类型").Field(field.New("type", t.dt)).Build()
	}

	return
}

func (t *Table) dialect() (smallint, bigint, datetime, varchar string, err error) {
	switch t.dt {
	case db.TypeMySQL:
		smallint = "SMALLINT"
		bigint = "BIGINT"
		datetime = "DATETIME"
		varchar = "VARCHAR(1024)"
	case db.TypePostgres:
		smallint = "SMALLINT"
		bigint = "BIGINT"
		datetime = "TIMESTAMP"
		varchar = "VARCHAR(1024)"
	case db.TypeSQLite:
		smallint = "INTEGER"
		bigint = "INTEGER"
		datetime = "NUMERIC"
		varchar = "TEXT(1024)"
	case db.TypeOracle:
		smallint = "SMALLINT"
		bigint = "BIGINT"
		datetime = "DATETIME"
	default:
		err = exception.New().Message("不被支持的数据库类型").Field(field.New("type", t.dt)).Build()
	}

	return
}
