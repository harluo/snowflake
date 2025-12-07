package core

import (
	"github.com/goexl/id"
	"github.com/goexl/snowflake"
	"github.com/harluo/snowflake/internal/config"
)

type Generator = id.Generator

func newGenerator(config *config.Snowflake) Generator {
	return snowflake.New().
		Machine(config.Machine).
		Started(config.Started).
		Build()
}
