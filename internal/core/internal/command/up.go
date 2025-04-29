package command

import (
	"context"

	"github.com/goexl/log"
	"github.com/pangum/migration/internal/core/internal/command/internal"
	"github.com/pangum/migration/internal/kernel"
)

type Up struct {
	migrations []kernel.Migration
	logger     log.Logger
}

func newUp(get internal.Get) *Up {
	return &Up{
		migrations: get.Migrations,
		logger:     get.Logger,
	}
}

func (u *Up) Name() string {
	return "up"
}

func (u *Up) Aliases() []string {
	return []string{
		"u",
		"upgrade",
	}
}

func (u *Up) Usage() string {
	return "升级"
}

func (u *Up) Run(ctx context.Context) (err error) {
	return
}

func (u *Up) Description() string {
	return `升级`
}
