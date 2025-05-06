package checker

import (
	"context"
)

type Downgrader interface {
	// Downgrade 降级时调用
	Downgrade(context.Context) error
}
