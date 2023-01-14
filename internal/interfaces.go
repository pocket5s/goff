package internal

import (
	"context"
)

type repository interface {
	save(ctx context.Context, flag *goff.Flag) error
	shutdown()
}
