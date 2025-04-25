package config

import (
	"github.com/harluo/config"
)

type Id struct {
	// 雪花算法配置
	Snowflake *Snowflake `json:"snowflake,omitempty" validate:"required"`
}

func newId(config *config.Getter) (id *Id, err error) {
	id = new(Id)
	err = config.Get(&struct {
		Id *Id `json:"id,omitempty" validate:"required"`
	}{
		Id: id,
	})

	return
}
