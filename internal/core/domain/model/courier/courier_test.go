package courier

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"delivery/internal/core/domain/kernel"
	"delivery/internal/core/domain/model/order"
	"delivery/internal/pkg/errs"
)

func Test_NewCourierCorrectCreation(t *testing.T) {
	location, err := kernel.NewLocation(1, 1)
	assert.NoError(t, err)

	courier, err := NewCourier("John", 1, location)
	assert.NoError(t, err)

	assert.NotEmpty(t, courier)
	assert.Equal(t, "John", courier.Name())
	assert.Equal(t, 1, courier.Speed())
}

func Test_NewCourierErrIncorrectParams(t *testing.T) {
	location, err := kernel.NewLocation(1, 1)
	assert.NoError(t, err)

	tests := map[string]struct {
		name     string
		speed    int
		location kernel.Location
		expected error
	}{
		"Name: '', speed: 1, location: not empty": {
			name:     "",
			speed:    1,
			location: location,
			expected: errs.NewValueIsInvalidError("name"),
		},
		"Name: John, speed: -2, location: not empty": {
			name:     "John",
			speed:    -2,
			location: location,
			expected: errs.NewValueIsInvalidError("speed"),
		},
		"Name: John, speed: 1, location: empty": {
			name:     "John",
			speed:    1,
			location: kernel.Location{},
			expected: errs.NewValueIsInvalidError("location"),
		},
	}
	for n, test := range tests {
		t.Run(n, func(t *testing.T) {
			_, err := NewCourier(test.name, test.speed, test.location)

			if err.Error() != test.expected.Error() {
				t.Errorf("expected %v, got %v", test.expected, err)
			}
		})
	}
}

func Test_CourierAddStoragePlaceCorrect(t *testing.T) {
	location, err := kernel.NewLocation(1, 1)
	assert.NoError(t, err)

	courier, err := NewCourier("John", 1, location)
	assert.NoError(t, err)

	placeName := "Place2"
	placeVolume := 15

	err = courier.AddStoragePlace(placeName, placeVolume)
	assert.NoError(t, err)

	storagePlaces := courier.StoragePlaces()
	for _, place := range storagePlaces {
		if place.Name() != placeName {
			continue
		}
		assert.Equal(t, placeVolume, place.TotalVolume())
	}
}

func Test_CourierCanTakeOrderCorrect(t *testing.T) {
	location, err := kernel.NewLocation(1, 1)
	assert.NoError(t, err)

	courier, err := NewCourier("John", 1, location)
	assert.NoError(t, err)

	order1, err := order.NewOrder(uuid.New(), location, 5)
	assert.NoError(t, err)

	order2, err := order.NewOrder(uuid.New(), location, 100)
	assert.NoError(t, err)

	tests := map[string]struct {
		order    *order.Order
		expected bool
	}{
		"Courier can take order": {
			order:    order1,
			expected: true,
		},
		"Courier cannot take order": {
			order:    order2,
			expected: false,
		},
	}
	for n, test := range tests {
		t.Run(n, func(t *testing.T) {
			res, err := courier.CanTakeOrder(test.order)
			assert.Equal(t, res, test.expected)
			assert.NoError(t, err)
		})
	}
}

func Test_CourierTakeOrderCorrect(t *testing.T) {
	location, err := kernel.NewLocation(1, 1)
	assert.NoError(t, err)

	courier, err := NewCourier("John", 1, location)
	assert.NoError(t, err)

	order1, err := order.NewOrder(uuid.New(), location, 5)
	assert.NoError(t, err)

	err = courier.TakeOrder(order1)
	assert.NoError(t, err)
}

func Test_CourierTakeOrderError(t *testing.T) {
	location, err := kernel.NewLocation(1, 1)
	assert.NoError(t, err)

	courier, err := NewCourier("John", 1, location)
	assert.NoError(t, err)

	order1, err := order.NewOrder(uuid.New(), location, 20)
	assert.NoError(t, err)

	err = courier.TakeOrder(order1)
	assert.EqualError(t, err, ErrCourierCannotTakeOrder.Error())
}

func Test_CourierCanCalculateTimeToLocation(t *testing.T) {
	// Изначальная точка курьера: [1,1]
	// Целевая точка: [5,10]
	// Количество шагов, необходимое курьеру: 13 (4 по горизонтали и 9 по вертикали)
	// Скорость транспорта (велосипедиста): 2 шага в 1 такт
	// Время подлета: 13/2 = 6.5 тактов потребуется курьеру, чтобы доставить заказ

	// Arrange
	courier, err := NewCourier("Велосипедист", 2, kernel.MinLocation())
	assert.NoError(t, err)
	target, err := kernel.NewLocation(5, 10)
	assert.NoError(t, err)

	// Act
	time, err := courier.CalculateTimeToLocation(target)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 6.5, time)
}
