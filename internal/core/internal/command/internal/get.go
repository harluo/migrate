package internal

import (
	"github.com/goexl/log"
	"github.com/harluo/di"
	"github.com/harluo/migrate/internal/core/internal/command/internal/argument"
	"github.com/harluo/migrate/internal/internal/db"
	"github.com/harluo/migrate/internal/kernel"
)

type Get struct {
	di.Get

	Table      *db.Table
	Upgrade    *db.Upgrade
	Downgrade  *db.Downgrade
	Migrations []kernel.Migration `group:"migrations"`

	Id      *argument.Id
	Pattern *argument.Pattern
	Logger  log.Logger
}
