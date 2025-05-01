package command

import (
	"context"

	"github.com/goexl/log"
	"github.com/harluo/migrate/internal/core/internal/command/internal"
	"github.com/harluo/migrate/internal/kernel"
)

type Downgrade struct {
	migrations []kernel.Migration
	logger     log.Logger
}

func newDown(get internal.Get) *Downgrade {
	return &Downgrade{
		migrations: get.Migrations,
		logger:     get.Logger,
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
	return "降级"
}

func (d *Downgrade) Run(ctx context.Context) (err error) {
	return
}

func (d *Downgrade) Description() string {
	return `降级`
}
