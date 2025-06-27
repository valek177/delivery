package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"delivery/internal/core/domain/kernel"
	courierLib "delivery/internal/core/domain/model/courier"
	"delivery/internal/core/domain/model/order"
	"delivery/internal/pkg/errs"
)

func Test_OrderDispatcherCorrectDispatch(t *testing.T) {
	dispatcher := NewOrderDispatcher()

	locationOrder, err := kernel.NewLocation(4, 5)
	assert.NoError(t, err)

	locationCourier1, err := kernel.NewLocation(5, 7)
	assert.NoError(t, err)
	locationCourier2, err := kernel.NewLocation(1, 1)
	assert.NoError(t, err)

	order1, err := order.NewOrder(uuid.New(), locationOrder, 5)
	assert.NoError(t, err)

	courier1, err := courierLib.NewCourier("John", 1, locationCourier1)
	assert.NoError(t, err)
	courier2, err := courierLib.NewCourier("Peter", 1, locationCourier2)
	assert.NoError(t, err)

	bestCourier, err := dispatcher.Dispatch(order1, []*courierLib.Courier{courier1, courier2})
	assert.NoError(t, err)

	assert.NotEmpty(t, bestCourier)
	assert.Equal(t, "John", bestCourier.Name())
}

func Test_OrderDispatcherDispatchNoSuitableCourier(t *testing.T) {
	dispatcher := NewOrderDispatcher()

	locationOrder, err := kernel.NewLocation(4, 5)
	assert.NoError(t, err)

	locationCourier, err := kernel.NewLocation(5, 7)
	assert.NoError(t, err)

	order1, err := order.NewOrder(uuid.New(), locationOrder, 20)
	assert.NoError(t, err)

	courier, err := courierLib.NewCourier("John", 1, locationCourier)
	assert.NoError(t, err)

	bestCourier, err := dispatcher.Dispatch(order1, []*courierLib.Courier{courier})
	assert.Error(t, err)
	assert.EqualError(t, err, "no suitable couriers")

	assert.Empty(t, bestCourier)
}

func Test_OrderDispatcherDispatchErrors(t *testing.T) {
	dispatcher := NewOrderDispatcher()

	locationOrder, err := kernel.NewLocation(4, 5)
	assert.NoError(t, err)

	locationCourier, err := kernel.NewLocation(5, 7)
	assert.NoError(t, err)

	order1, err := order.NewOrder(uuid.New(), locationOrder, 7)
	assert.NoError(t, err)

	courier1, err := courierLib.NewCourier("John", 1, locationCourier)
	assert.NoError(t, err)

	tests := map[string]struct {
		order       *order.Order
		couriers    []*courierLib.Courier
		expectedErr error
	}{
		"Order is empty": {
			order:       nil,
			couriers:    []*courierLib.Courier{courier1},
			expectedErr: errs.NewValueIsInvalidError("order"),
		},
		"Couriers are empty": {
			order:       order1,
			couriers:    []*courierLib.Courier{},
			expectedErr: errs.NewValueIsRequiredError("couriers"),
		},
	}
	for n, test := range tests {
		t.Run(n, func(t *testing.T) {
			_, err := dispatcher.Dispatch(test.order, test.couriers)

			if err.Error() != test.expectedErr.Error() {
				t.Errorf("expected %v, got %v", test.expectedErr, err)
			}
		})
	}
}
