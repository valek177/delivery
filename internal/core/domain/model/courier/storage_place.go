package courier

import (
	"fmt"

	"delivery/internal/pkg/ddd"
	"delivery/internal/pkg/errs"

	"github.com/google/uuid"
)

var (
	ErrStoragePlaceIsOccupied      = fmt.Errorf("storage place already occupied")
	ErrOrderVolumeExceedsTotal     = fmt.Errorf("order volume exceeds storage place total volume")
	ErrCannotStoreInStoragePlace   = fmt.Errorf("cannot store order in storage place")
	ErrOrderNotFoundInStoragePlace = fmt.Errorf("cannot find order in storage place")
)

type StoragePlace struct {
	baseEntity *ddd.BaseEntity[uuid.UUID]

	id          uuid.UUID
	name        string
	totalVolume int
	orderID     *uuid.UUID
}

func NewStoragePlace(name string, totalVolume int) (*StoragePlace, error) {
	if name == "" {
		return nil, errs.NewValueIsRequiredError("name")
	}
	if totalVolume <= 0 {
		return nil, errs.NewValueIsRequiredError("totalVolume")
	}

	return &StoragePlace{
		id:          uuid.New(),
		name:        name,
		totalVolume: totalVolume,
	}, nil
}

func RestoreStoragePlace(id uuid.UUID, name string, totalVolume int,
	orderID *uuid.UUID,
) *StoragePlace {
	return &StoragePlace{
		baseEntity:  ddd.NewBaseEntity(id),
		name:        name,
		totalVolume: totalVolume,
		orderID:     orderID,
	}
}

func (s *StoragePlace) Equals(other *StoragePlace) bool {
	return s.id == other.id
}

func (s *StoragePlace) ID() uuid.UUID {
	return s.id
}

func (s *StoragePlace) Name() string {
	return s.name
}

func (s *StoragePlace) TotalVolume() int {
	return s.totalVolume
}

func (s *StoragePlace) OrderID() *uuid.UUID {
	return s.orderID
}

func (s *StoragePlace) CanStore(volume int) (bool, error) {
	if volume <= 0 {
		return false, errs.NewValueIsInvalidError("volume")
	}
	if volume > s.TotalVolume() {
		return false, ErrOrderVolumeExceedsTotal
	}
	if s.isOccupied() {
		return false, ErrStoragePlaceIsOccupied
	}

	return true, nil
}

func (s *StoragePlace) Store(orderID uuid.UUID, volume int) error {
	if orderID == uuid.Nil {
		return errs.NewValueIsRequiredError("orderID")
	}
	if volume <= 0 {
		return errs.NewValueIsInvalidError("volume")
	}

	canStore, err := s.CanStore(volume)
	if err != nil {
		return err
	}

	if canStore {
		s.orderID = &orderID
	} else {
		return ErrCannotStoreInStoragePlace
	}

	return nil
}

func (s *StoragePlace) Clear(orderID uuid.UUID) error {
	if orderID == uuid.Nil {
		return errs.NewValueIsRequiredError("orderID")
	}
	if s.orderID == nil || *s.orderID != orderID {
		return ErrOrderNotFoundInStoragePlace
	}

	s.orderID = nil

	return nil
}

func (s *StoragePlace) isOccupied() bool {
	return s.OrderID() != nil
}
