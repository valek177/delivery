package queries

import "github.com/google/uuid"

type GetNotCompletedOrdersResponse struct {
	Orders []NotCompletedOrdersResponse
}

type NotCompletedOrdersResponse struct {
	ID       uuid.UUID        `gorm:"type:uuid;primaryKey"`
	Location LocationResponse `gorm:"embedded;embeddedPrefix:location_"`
}

func (NotCompletedOrdersResponse) TableName() string {
	return "orders"
}
