package command

import (
	"context"

	"github.com/goexl/log"
	"github.com/harluo/migrate/internal/core/internal/command/internal"
	"github.com/harluo/migrate/internal/kernel"
)

type Upgrade struct {
	migrations []kernel.Migration
	logger     log.Logger
}

func newUpgrade(get internal.Get) *Upgrade {
	return &Upgrade{
		migrations: get.Migrations,
		logger:     get.Logger,
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

func (u *Upgrade) Run(ctx context.Context) (err error) {
	return
}

func (u *Upgrade) Description() string {
	return `升级`
}
