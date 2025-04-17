package config

import (
	"time"
)

type Snowflake struct {
	// 跳过数量
	Skip uint64 `json:"skip,omitempty"`
	// 时间点，有此配置后，生成的标识时间都为此配置时间
	Epoch time.Time `json:"epoch,omitempty"`
	// 集群标识
	Node uint64 `default:"1" json:"node,omitempty" validate:"required,min=1,max=32"`
	// 工作节点标识
	Worker uint64 `default:"1" json:"worker,omitempty" validate:"required,min=1,max=32"`
	// 顺序号
	Sequence uint64 `default:"1" json:"sequence,omitempty" validate:"required,min=0,max=4096"`
}

func newSnowflake(config *Id) *Snowflake {
	return config.Snowflake
}
