package core

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newGenerator,
	).Build().Apply()
}
