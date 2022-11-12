package client

import (
	"github.com/vleukhin/GophKeeper/internal/client/core"
	"sync"
)

var (
	app  *core.Core //nolint:gochecknoglobals
	once sync.Once  //nolint:gochecknoglobals
)

func GetApp() *core.Core {
	once.Do(func() {
		app = &core.Core{}
	})

	return app
}
