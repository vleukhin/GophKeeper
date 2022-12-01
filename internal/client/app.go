package client

import (
	"sync"

	"github.com/vleukhin/GophKeeper/internal/client/core"
)

var (
	app  *core.Core //nolint:gochecknoglobals
	once sync.Once  //nolint:gochecknoglobals
)

func GetApp() core.GophKeeperClient {
	once.Do(func() {
		app = &core.Core{}
	})

	return app
}
