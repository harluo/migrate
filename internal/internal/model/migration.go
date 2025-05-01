package model

import (
	"time"
)

type Migration struct {
	Id uint64 `json:"id,omitempty"`
	// 版本号
	Version uint16 `json:"version,omitempty"`
	// 描述
	Description string `json:"description,omitempty"`

	// 创建时间
	Created time.Time `json:"created,omitempty"`
}

func (*Migration) TableComment() string {
	return "数据迁移记录"
}
