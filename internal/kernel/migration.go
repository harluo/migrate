package kernel

import (
	"context"
)

type Migration interface {
	// Id 标识
	Id() uint64

	// Upgrade 升级时调用
	Upgrade(context.Context) error

	// Downgrade 降级时调用
	Downgrade(context.Context) error

	// Description 升级描述
	Description() string
}
