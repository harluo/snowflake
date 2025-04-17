package config

import (
	"github.com/harluo/boot"
)

type Id struct {
	// 雪花算法配置
	Snowflake *Snowflake `json:"snowflake,omitempty" validate:"required"`
}

func newId(config *boot.Config) (id *Id, err error) {
	id = new(Id)
	err = config.Build().Get(&struct {
		Id *Id `json:"id,omitempty" validate:"required"`
	}{
		Id: id,
	})

	return
}
