package courier

import (
	"errors"
	"fmt"
	"math"

	"github.com/google/uuid"

	"delivery/internal/core/domain/kernel"
	"delivery/internal/core/domain/model/order"
	"delivery/internal/pkg/ddd"
	"delivery/internal/pkg/errs"
)

var ErrCourierCannotTakeOrder = errors.New("courier cannot take order")

const (
	courierDefaultStoragePlaceName   = "Bag"
	courierDefaultStoragePlaceVolume = 10
)

type Courier struct {
	baseAggregate *ddd.BaseAggregate[uuid.UUID]

	name          string
	speed         int
	location      kernel.Location
	storagePlaces []*StoragePlace
}

func NewCourier(name string, speed int, location kernel.Location) (*Courier, error) {
	if name == "" {
		return nil, errs.NewValueIsInvalidError("name")
	}
	if speed <= 0 {
		return nil, errs.NewValueIsInvalidError("speed")
	}
	if location.IsEmpty() {
		return nil, errs.NewValueIsInvalidError("location")
	}

	storagePlace, err := NewStoragePlace(courierDefaultStoragePlaceName,
		courierDefaultStoragePlaceVolume)
	if err != nil {
		return nil,
			fmt.Errorf("cannot create courier: cannot create storage place: %v", err)
	}

	return &Courier{
		baseAggregate: ddd.NewBaseAggregate(uuid.New()),
		name:          name,
		speed:         speed,
		location:      location,
		storagePlaces: []*StoragePlace{
			storagePlace,
		},
	}, nil
}

func RestoreCourier(id uuid.UUID, name string, speed int, location kernel.Location,
	storagePlaces []*StoragePlace,
) *Courier {
	return &Courier{
		baseAggregate: ddd.NewBaseAggregate(id),
		name:          name,
		speed:         speed,
		location:      location,
		storagePlaces: storagePlaces,
	}
}

func (c *Courier) Equals(other *Courier) bool {
	if other == nil {
		return false
	}

	return c.baseAggregate.Equal(other.baseAggregate)
}

func (c *Courier) ClearDomainEvents() {
	c.baseAggregate.ClearDomainEvents()
}

func (c *Courier) GetDomainEvents() []ddd.DomainEvent {
	return c.baseAggregate.GetDomainEvents()
}

func (c *Courier) RaiseDomainEvent(event ddd.DomainEvent) {
	c.baseAggregate.RaiseDomainEvent(event)
}

func (c *Courier) ID() uuid.UUID {
	return c.baseAggregate.ID()
}

func (c *Courier) Name() string {
	return c.name
}

func (c *Courier) Speed() int {
	return c.speed
}

func (c *Courier) Location() kernel.Location {
	return c.location
}

func (c *Courier) StoragePlaces() []*StoragePlace {
	return c.storagePlaces
}

func (c *Courier) AddStoragePlace(name string, volume int) error {
	newStoragePlace, err := NewStoragePlace(name, volume)
	if err != nil {
		return fmt.Errorf("cannot add storage place: %v", err)
	}

	c.storagePlaces = append(c.storagePlaces, newStoragePlace)

	return nil
}

func (c *Courier) CanTakeOrder(order *order.Order) (bool, error) {
	if order == nil {
		return false, errs.NewValueIsInvalidError("order")
	}

	for _, place := range c.storagePlaces {
		canStore, err := place.CanStore(order.Volume())
		if err != nil {
			continue
		}
		if canStore {
			return true, nil
		}
	}

	return false, nil
}

func (c *Courier) TakeOrder(order *order.Order) error {
	if order == nil {
		return errs.NewValueIsInvalidError("order")
	}

	canTake, err := c.CanTakeOrder(order)
	if err != nil {
		return err
	}

	if !canTake {
		return ErrCourierCannotTakeOrder
	}

	for _, place := range c.storagePlaces {
		canStore, err := place.CanStore(order.Volume())
		if err != nil {
			continue
		}
		if canStore {
			err = place.Store(order.ID(), order.Volume())
			if err != nil {
				return err
			}
			return nil
		}
	}

	return ErrCourierCannotTakeOrder
}

func (c *Courier) CompleteOrder(order *order.Order) error {
	if order == nil {
		return errs.NewValueIsRequiredError("order")
	}

	storagePlace, err := c.findStoragePlaceByOrderID(order.ID())
	if err != nil {
		return fmt.Errorf("cannot complete order: %v", err)
	}

	err = storagePlace.Clear(order.ID())
	if err != nil {
		return err
	}

	return nil
}

func (c *Courier) CalculateTimeToLocation(target kernel.Location) (float64, error) {
	if target.IsEmpty() {
		return 0, errs.NewValueIsRequiredError("target")
	}

	distance, err := c.location.DistanceTo(target)
	if err != nil {
		return 0, err
	}

	time := float64(distance) / float64(c.speed)
	return time, err
}

func (c *Courier) Move(target kernel.Location) error {
	if target.IsEmpty() {
		return errs.NewValueIsRequiredError("target")
	}

	dx := float64(target.X() - c.location.X())
	dy := float64(target.Y() - c.location.Y())
	remainingRange := float64(c.speed)

	if math.Abs(dx) > remainingRange {
		dx = math.Copysign(remainingRange, dx)
	}
	remainingRange -= math.Abs(dx)

	if math.Abs(dy) > remainingRange {
		dy = math.Copysign(remainingRange, dy)
	}

	newX := c.location.X() + int(dx)
	newY := c.location.Y() + int(dy)

	newLocation, err := kernel.NewLocation(newX, newY)
	if err != nil {
		return err
	}
	c.location = newLocation
	return nil
}

func (c *Courier) findStoragePlaceByOrderID(orderID uuid.UUID) (*StoragePlace, error) {
	if orderID == uuid.Nil {
		return nil, errs.NewValueIsInvalidError("orderID")
	}

	for _, place := range c.storagePlaces {
		if place.OrderID() != nil && orderID == *place.OrderID() {
			return place, nil
		}
	}

	return nil, fmt.Errorf("cannot find storage place by orderID: %s", orderID)
}
