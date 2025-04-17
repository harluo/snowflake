package config

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newId,        // 统一标识配置
		newSnowflake, // 配置
	).Build().Apply()
}
