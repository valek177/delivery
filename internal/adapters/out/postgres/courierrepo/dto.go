package courierrepo

import (
	"github.com/google/uuid"
)

type CourierDTO struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name          string
	Speed         int
	Location      LocationDTO        `gorm:"embedded;embeddedPrefix:location_"`
	StoragePlaces []*StoragePlaceDTO `gorm:"foreignKey:CourierID;constraint:OnDelete:CASCADE;"`
}

type StoragePlaceDTO struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	OrderID     *uuid.UUID `gorm:"type:uuid;"`
	Name        string
	TotalVolume int
	CourierID   uuid.UUID `gorm:"type:uuid;index"`
}

type LocationDTO struct {
	X int
	Y int
}

func (CourierDTO) TableName() string {
	return "couriers"
}

func (StoragePlaceDTO) TableName() string {
	return "storage_places"
}
