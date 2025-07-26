package uow

import (
	"gorm.io/gorm"

	"delivery/internal/core/ports"
	"delivery/internal/pkg/ddd"
)

type UnitOfWork interface {
	Tx() *gorm.DB
	Db() *gorm.DB
	InTx() bool
	Begin()
	Commit() error
	Track(agg ddd.AggregateRoot)
	CourierRepository() ports.CourierRepository
	OrderRepository() ports.OrderRepository
}
