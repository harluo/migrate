package core

import (
	"context"

	"github.com/pangum/migration/internal/core/internal/command"
)

type Initializer struct {
	up *command.Up
}

func newInitializer(up *command.Up) *Initializer {
	return &Initializer{
		up: up,
	}
}

func (i *Initializer) Initialize(ctx context.Context) error {
	return i.up.Run(ctx)
}
