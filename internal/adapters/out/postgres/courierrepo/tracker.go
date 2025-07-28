package courierrepo

import (
	"context"

	"gorm.io/gorm"

	"delivery/internal/pkg/ddd"
)

type Tracker interface {
	Tx() *gorm.DB
	Db() *gorm.DB
	InTx() bool
	Track(agg ddd.AggregateRoot)
	Begin(ctx context.Context)
	Commit(ctx context.Context) error
}
