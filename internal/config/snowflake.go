package config

import (
	"time"
)

type Snowflake struct {
	// 时间点，有此配置后，生成的标识时间都为此配置时间
	Started time.Time `json:"started,omitempty"`
	// 机器标识
	Machine uint32 `default:"1" json:"machine,omitempty" validate:"required,min=1,max=2147483648"`
}

func newSnowflake(config *Id) *Snowflake {
	return config.Snowflake
}
