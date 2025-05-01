package db

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newTable,
		newUpgrade,
		newDowngrade,
	).Build().Apply()
}
