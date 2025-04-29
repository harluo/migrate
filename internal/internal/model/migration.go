package model

import (
	"github.com/goexl/model"
)

type Migration struct {
	model.Base

	Key         string
	Name        string
	Description string
}

func (*Migration) TableComment() string {
	return "数据迁移记录"
}
