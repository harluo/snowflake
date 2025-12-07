package core

import (
	"github.com/goexl/id"
	"github.com/goexl/snowflake"
	"github.com/harluo/snowflake/internal/config"
)

type Generator = id.Generator

func newGenerator(config *config.Snowflake) (generator Generator, err error) {
	builder := snowflake.New().Machine(config.Machine)
	if nil != config.Started {
		if timestamp, te := config.Started.Time(); nil != te {
			err = te
		} else {
			builder.Started(timestamp)
		}
	}
	if nil == err {
		generator = builder.Build()
	}

	return
}
