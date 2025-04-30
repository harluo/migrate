package config

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newMigrate,
	).Build().Apply()
}
