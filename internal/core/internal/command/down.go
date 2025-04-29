package command

import (
	"context"

	"github.com/goexl/log"
	"github.com/harluo/migrate/internal/core/internal/command/internal"
	"github.com/harluo/migrate/internal/kernel"
)

type Down struct {
	migrations []kernel.Migration
	logger     log.Logger
}

func newDown(get internal.Get) *Down {
	return &Down{
		migrations: get.Migrations,
		logger:     get.Logger,
	}
}

func (d *Down) Name() string {
	return "down"
}

func (d *Down) Aliases() []string {
	return []string{
		"d",
		"downgrade",
	}
}

func (d *Down) Usage() string {
	return "降级"
}

func (d *Down) Run(ctx context.Context) (err error) {
	return
}

func (d *Down) Description() string {
	return `降级`
}
