package order

import (
	"errors"

	"github.com/google/uuid"

	"delivery/internal/core/domain/kernel"
	"delivery/internal/pkg/errs"
)

var (
	ErrCannotCompleteNotAssignedOrder   = errors.New("can not complete not assigned order")
	ErrCannotAssignAlreadyAssignedOrder = errors.New("can not assign already assigned order")
)

type Order struct {
	id        uuid.UUID
	courierID *uuid.UUID
	location  kernel.Location
	volume    int
	status    Status
}

func NewOrder(orderID uuid.UUID, location kernel.Location, volume int) (*Order, error) {
	if orderID == uuid.Nil {
		return nil, errs.NewValueIsRequiredError("orderID")
	}
	if location.IsEmpty() {
		return nil, errs.NewValueIsRequiredError("location")
	}
	if volume <= 0 {
		return nil, errs.NewValueIsRequiredError("volume")
	}
	return &Order{
		id:       orderID,
		location: location,
		volume:   volume,
		status:   StatusCreated,
	}, nil
}

func (o *Order) ID() uuid.UUID {
	return o.id
}

func (o *Order) CourierID() *uuid.UUID {
	return o.courierID
}

func (o *Order) Location() kernel.Location {
	return o.location
}

func (o *Order) Volume() int {
	return o.volume
}

func (o *Order) Status() Status {
	return o.status
}

func (o *Order) Assign(courierID uuid.UUID) error {
	if courierID == uuid.Nil {
		return errs.NewValueIsRequiredError("courierID")
	}
	if o.status != StatusCreated {
		return ErrCannotAssignAlreadyAssignedOrder
	}

	o.courierID = &courierID
	o.status = StatusAssigned
	return nil
}

func (o *Order) Complete() error {
	if o.status != StatusAssigned {
		return ErrCannotCompleteNotAssignedOrder
	}
	o.status = StatusCompleted
	return nil
}
