package command

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newUpgrade,
		newDown,
	).Build().Apply()
}
