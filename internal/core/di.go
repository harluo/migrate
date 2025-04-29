package core

import (
	"github.com/harluo/di"
	"github.com/harluo/migrate/internal/core/internal"
)

func init() {
	di.New().Instance().Put(
		newCommand,
		newInitializer,
		func(command *Command, initializer *Initializer) internal.Put {
			return internal.Put{
				Migrate:     command,
				Initializer: initializer,
			}
		},
	).Build().Apply()
}
