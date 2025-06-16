package courier

import (
	"testing"

	"delivery/internal/pkg/errs"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_NewStoragePlaceCorrectCreation(t *testing.T) {
	storagePlace, err := NewStoragePlace("place1", 10)

	assert.NoError(t, err)
	assert.NotEmpty(t, storagePlace)
	assert.Equal(t, "place1", storagePlace.Name())
	assert.Equal(t, 10, storagePlace.TotalVolume())
}

func Test_NewStoragePlaceErrIncorrectParams(t *testing.T) {
	tests := map[string]struct {
		name        string
		totalVolume int
		expected    error
	}{
		"Name: '', totalVolume: 100": {
			name:        "",
			totalVolume: 100,
			expected:    errs.NewValueIsRequiredError("name"),
		},
		"Name: place1, totalVolume: -5": {
			name:        "place1",
			totalVolume: -5,
			expected:    errs.NewValueIsRequiredError("totalVolume"),
		},
	}
	for n, test := range tests {
		t.Run(n, func(t *testing.T) {
			_, err := NewStoragePlace(test.name, test.totalVolume)

			if err.Error() != test.expected.Error() {
				t.Errorf("expected %v, got %v", test.expected, err)
			}
		})
	}
}

func Test_CanStoreReturnsTrue(t *testing.T) {
	storagePlace, err := NewStoragePlace("place1", 10)

	assert.NoError(t, err)

	res, err := storagePlace.CanStore(5)
	assert.NoError(t, err)
	assert.Equal(t, res, true)
}

func Test_CanStoreReturnsFalse(t *testing.T) {
	storagePlace, err := NewStoragePlace("place1", 10)

	assert.NoError(t, err)

	tests := map[string]struct {
		volume      int
		expected    bool
		expectedErr error
	}{
		"Volume: -5": {
			volume:      -5,
			expected:    false,
			expectedErr: errs.NewValueIsInvalidError("volume"),
		},
		"Volume: 1000": {
			volume:      1000,
			expected:    false,
			expectedErr: ErrOrderVolumeExceedsTotal,
		},
	}
	for n, test := range tests {
		t.Run(n, func(t *testing.T) {
			res, err := storagePlace.CanStore(test.volume)

			if res != test.expected {
				t.Errorf("expected %v, got %v", test.expected, res)
			}

			if err.Error() != test.expectedErr.Error() {
				t.Errorf("expected %v, got %v", test.expectedErr, err)
			}
		})
	}
}

func Test_CanStoreReturnsFalseOnOccupied(t *testing.T) {
	storagePlace, err := NewStoragePlace("place1", 10)
	assert.NoError(t, err)

	orderID := uuid.New()
	storagePlace.orderID = &orderID

	res, err := storagePlace.CanStore(5)
	assert.Equal(t, res, false)
	assert.Equal(t, err, ErrStoragePlaceIsOccupied)
}

func Test_StoreOk(t *testing.T) {
	storagePlace, err := NewStoragePlace("place1", 10)
	assert.NoError(t, err)

	orderID := uuid.New()
	err = storagePlace.Store(orderID, 5)
	assert.NoError(t, err)
}

func Test_StoreReturnsErrorOnIncorrectParams(t *testing.T) {
	storagePlace, err := NewStoragePlace("place1", 10)
	assert.NoError(t, err)

	tests := map[string]struct {
		orderID     uuid.UUID
		volume      int
		expectedErr error
	}{
		"OrderID: nil": {
			orderID:     uuid.Nil,
			volume:      5,
			expectedErr: errs.NewValueIsRequiredError("orderID"),
		},
		"OrderID is correct, volume: 0": {
			orderID:     uuid.New(),
			volume:      0,
			expectedErr: errs.NewValueIsInvalidError("volume"),
		},
	}
	for n, test := range tests {
		t.Run(n, func(t *testing.T) {
			err := storagePlace.Store(test.orderID, test.volume)

			if err.Error() != test.expectedErr.Error() {
				t.Errorf("expected %v, got %v", test.expectedErr, err)
			}
		})
	}
}

func Test_ClearOk(t *testing.T) {
	storagePlace, err := NewStoragePlace("place1", 10)
	assert.NoError(t, err)

	orderID := uuid.New()

	err = storagePlace.Store(orderID, 5)
	assert.NoError(t, err)

	err = storagePlace.Clear(orderID)
	assert.NoError(t, err)
}

func Test_ClearReturnsErrorOnIncorrectParams(t *testing.T) {
	storagePlace, err := NewStoragePlace("place1", 10)
	assert.NoError(t, err)

	orderID := uuid.New()

	err = storagePlace.Store(orderID, 5)
	assert.NoError(t, err)

	tests := map[string]struct {
		orderID     uuid.UUID
		expectedErr error
	}{
		"OrderID: nil": {
			orderID:     uuid.Nil,
			expectedErr: errs.NewValueIsRequiredError("orderID"),
		},
		"OrderID: occupied by another orderID": {
			orderID:     uuid.New(),
			expectedErr: ErrOrderNotFoundInStoragePlace,
		},
	}
	for n, test := range tests {
		t.Run(n, func(t *testing.T) {
			err := storagePlace.Clear(test.orderID)

			if err.Error() != test.expectedErr.Error() {
				t.Errorf("expected %v, got %v", test.expectedErr, err)
			}
		})
	}
}
