package queries

import (
	"gorm.io/gorm"

	"delivery/internal/pkg/errs"
)

type GetCouriersQueryHandler interface {
	Handle(GetCouriersQuery) (GetCouriersResponse, error)
}

type getCouriersQueryHandler struct {
	db *gorm.DB
}

func NewGetCouriersQueryHandler(
	db *gorm.DB,
) (GetCouriersQueryHandler, error) {
	if db == nil {
		return &getCouriersQueryHandler{}, errs.NewValueIsRequiredError("db")
	}

	return &getCouriersQueryHandler{
		db: db,
	}, nil
}

func (q *getCouriersQueryHandler) Handle(query GetCouriersQuery) (
	GetCouriersResponse, error,
) {
	var couriers []CourierResponse
	result := q.db.Raw("SELECT id, name, location_x, location_y FROM couriers").Scan(&couriers)

	if result.Error != nil {
		return GetCouriersResponse{}, result.Error
	}

	return GetCouriersResponse{Couriers: couriers}, nil
}
