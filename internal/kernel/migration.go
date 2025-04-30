package kernel

import (
	"context"
)

type Migration interface {
	// Id 标识
	Id() uint64

	// Up 升级时调用
	Up(context.Context) error

	// Down 降级时调用
	Down(context.Context) error

	// Description 升级描述
	Description() string
}
