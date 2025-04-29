package internal

import (
	"github.com/goexl/log"
	"github.com/harluo/di"
	"github.com/pangum/migration/internal/kernel"
)

type Get struct {
	di.Get

	Migrations []kernel.Migration `group:"migrations"`
	Logger     log.Logger
}
