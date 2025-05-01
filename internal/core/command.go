package core

import (
	"context"

	"github.com/goexl/log"
	"github.com/harluo/boot"
	"github.com/harluo/migrate/internal/core/internal/command"
)

type Command struct {
	up     *command.Upgrade
	down   *command.Downgrade
	logger log.Logger
}

func newCommand(up *command.Upgrade, down *command.Downgrade, logger log.Logger) *Command {
	return &Command{
		up:     up,
		down:   down,
		logger: logger,
	}
}

func (c *Command) Run(ctx context.Context) error {
	return c.up.Run(ctx) // 默认执行升级命令
}

func (*Command) Name() string {
	return "migrate"
}

func (*Command) Aliases() []string {
	return []string{
		"m",
		"mgr",
	}
}

func (*Command) Usage() string {
	return "数据迁移"
}

func (*Command) Description() string {
	return `执行数据迁移操作`
}

func (c *Command) Subcommands() []boot.Command {
	return []boot.Command{
		c.up,
		c.down,
	}
}
