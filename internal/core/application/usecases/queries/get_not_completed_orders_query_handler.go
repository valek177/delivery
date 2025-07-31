package queries

import (
	"gorm.io/gorm"

	"delivery/internal/pkg/errs"
)

type GetNotCompletedOrdersQueryHandler interface {
	Handle(GetNotCompletedOrdersQuery) (GetNotCompletedOrdersResponse, error)
}

type getNotCompletedOrdersQueryHandler struct {
	db *gorm.DB
}

func NewGetNotCompletedOrdersQueryHandler(
	db *gorm.DB,
) (GetNotCompletedOrdersQueryHandler, error) {
	if db == nil {
		return &getNotCompletedOrdersQueryHandler{}, errs.NewValueIsRequiredError("db")
	}

	return &getNotCompletedOrdersQueryHandler{
		db: db,
	}, nil
}

func (q *getNotCompletedOrdersQueryHandler) Handle(query GetNotCompletedOrdersQuery) (
	GetNotCompletedOrdersResponse, error,
) {
	var orders []NotCompletedOrdersResponse
	result := q.db.Raw("SELECT id, location_x, location_y FROM orders WHERE status != ?").Scan(&orders)

	if result.Error != nil {
		return GetNotCompletedOrdersResponse{}, result.Error
	}

	return GetNotCompletedOrdersResponse{Orders: orders}, nil
}
