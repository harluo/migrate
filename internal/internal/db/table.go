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
	if smallint, bigint, datetime, varchar, de := t.dialect(); nil != de {
		err = de
	} else if exists, ee := t.exists(); nil != ee {
		err = ee
	} else if !exists {
		_, err = t.db.ExecContext(ctx, fmt.Sprintf(`CREATE TABLE %s (
	id %s PRIMARY KEY COMMENT '标识',

	version %s NOT NULL DEFAULT 1 COMMENT '版本号',
	description %s NOT NULL DEFAULT '' COMMENT '本次数据迁移描述信息',

	created %s NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '本次数据迁移创建时间',
	updated %s NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '本次数据迁移修改时间'
) COMMENT '数据迁移记录'`, t.config.Table, bigint, smallint, varchar, datetime, datetime))
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
		sql = `SELECT EXISTS(SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = $1)`
	case db.TypeSQLite:
		sql = `SELECT "name" FROM sqlite_master WHERE type='table' AND name = ?`
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
