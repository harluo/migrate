package config

import (
	"github.com/harluo/config"
)

type Migrate struct {
	// 是否开启功能
	Enabled bool `json:"enabled,omitempty"`
	// 表名
	Table string `default:"migration" json:"table,omitempty"`
}

func newMigrate(config config.Getter) (migrate *Migrate, err error) {
	migrate = new(Migrate)
	err = config.Get(&struct {
		Migrate *Migrate `json:"migrate,omitempty" validate:"required"`
	}{
		Migrate: migrate,
	})

	return
}
