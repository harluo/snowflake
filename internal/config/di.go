package config

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newId,        // 统一标识配置
		newSnowflake, // 配置
	).Build().Apply()
}
