package core

import (
	"context"

	"github.com/goexl/log"
	"github.com/harluo/boot"
	"github.com/harluo/migrate/internal/core/internal/command"
	"github.com/harluo/migrate/internal/internal/db"
)

type Command struct {
	upgrade   *command.Upgrade
	downgrade *command.Downgrade
	table     *db.Table
	logger    log.Logger
}

func newCommand(upgrade *command.Upgrade, downgrade *command.Downgrade, table *db.Table, logger log.Logger) *Command {
	return &Command{
		upgrade:   upgrade,
		downgrade: downgrade,
		table:     table,
		logger:    logger,
	}
}

func (c *Command) Run(ctx context.Context) (err error) {
	if ce := c.table.Create(ctx); nil != ce {
		err = ce
	} else {
		err = c.upgrade.Run(ctx) // 默认执行升级命令
	}

	return
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
		c.upgrade,
		c.downgrade,
	}
}
