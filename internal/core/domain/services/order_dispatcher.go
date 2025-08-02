package services

import (
	"fmt"
	"math"

	"delivery/internal/core/domain/model/courier"
	orderModel "delivery/internal/core/domain/model/order"
	"delivery/internal/pkg/errs"
)

type OrderDispatcher interface {
	Dispatch(order *orderModel.Order, couriers []*courier.Courier) (*courier.Courier, error)
}

type orderDispatcher struct{}

func NewOrderDispatcher() OrderDispatcher {
	return &orderDispatcher{}
}

func (d *orderDispatcher) Dispatch(order *orderModel.Order, couriers []*courier.Courier,
) (*courier.Courier, error) {
	if order == nil {
		return nil, errs.NewValueIsInvalidError("order")
	}

	if len(couriers) == 0 {
		return nil, errs.NewValueIsRequiredError("couriers")
	}

	if order.Status() != orderModel.StatusCreated {
		return nil, fmt.Errorf("order status is not created")
	}

	courier, err := d.bestCourier(order, couriers)
	if err != nil {
		return nil, err
	}

	fmt.Println("order", order.ID())

	err = courier.TakeOrder(order)
	if err != nil {
		return nil, err
	}

	err = order.Assign(courier.ID())
	if err != nil {
		return nil, err
	}

	return courier, nil
}

func (d *orderDispatcher) bestCourier(order *orderModel.Order, couriers []*courier.Courier,
) (*courier.Courier, error) {
	minTime := math.MaxFloat64
	var bestCourier *courier.Courier

	fmt.Println("couriers", couriers)

	for _, courier := range couriers {
		canTake, err := courier.CanTakeOrder(order)
		if err != nil {
			return nil, err
		}
		fmt.Println("can take", canTake)

		if !canTake {
			continue
		}

		time, err := courier.CalculateTimeToLocation(order.Location())
		fmt.Println("time ", time)
		if err != nil {
			return nil, err
		}
		if time < minTime {
			minTime = time
			bestCourier = courier
		}
	}
	fmt.Println("best couriers", bestCourier)

	if bestCourier == nil {
		return nil, fmt.Errorf("no suitable couriers")
	}

	return bestCourier, nil
}
