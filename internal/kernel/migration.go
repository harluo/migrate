package kernel

import (
	"context"
)

type Migration interface {
	// Name 升级名称
	Name() string

	// Up 升级时调用
	Up(context.Context) error

	// Down 降级时调用
	Down(context.Context) error
}
