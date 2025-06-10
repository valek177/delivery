package kernel

import (
	"testing"

	"delivery/internal/pkg/errs"

	"github.com/stretchr/testify/assert"
)

func Test_LocationBeCorrectWhenParamsAreCorrectOnCreated(t *testing.T) {
	location, err := NewLocation(4, 5)

	assert.NoError(t, err)
	assert.NotEmpty(t, location)
	assert.Equal(t, Location{4, 5, true}, location)
}

func Test_LocationReturnErrorWhenParamsAreIncorrectOnCreated(t *testing.T) {
	tests := map[string]struct {
		x        int
		y        int
		expected error
	}{
		"x: -1, y: 2": {
			x:        -1,
			y:        2,
			expected: errs.NewValueIsOutOfRangeError("x", -1, minX, maxX),
		},
		"x: 2, y: -100": {
			x:        2,
			y:        -100,
			expected: errs.NewValueIsOutOfRangeError("y", -100, minY, maxY),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := NewLocation(test.x, test.y)

			if err.Error() != test.expected.Error() {
				t.Errorf("expected %v, got %v", test.expected, err)
			}
		})
	}
}

func Test_LocationNewRandomReturnsCorrectLocation(t *testing.T) {
	for range 10 {
		location := NewRandomLocation()

		assert.False(t, location.IsEmpty())
		assert.Less(t, location.X(), maxX+1)
		assert.Greater(t, location.X(), minX-1)
		assert.Less(t, location.Y(), maxY+1)
		assert.Greater(t, location.Y(), minY-1)
	}
}

func Test_LocationReturnTrueIfLocationsAreEqual(t *testing.T) {
	location1, _ := NewLocation(2, 3)
	location2, _ := NewLocation(2, 3)

	assert.True(t, location1.Equals(location2))
}

func Test_LocationReturnTrueIfLocationIsEmpty(t *testing.T) {
	var location Location

	assert.True(t, location.IsEmpty())
}

func Test_LocationReturnCorrectDistanceBetweenLocations(t *testing.T) {
	location1, _ := NewLocation(2, 3)
	location2, _ := NewLocation(8, 7)

	dist, err := location1.DistanceTo(location2)
	assert.NoError(t, err)

	assert.Equal(t, 10, dist)
}

func Test_LocationReturnErrorIfTargetIsEmpty(t *testing.T) {
	location1, _ := NewLocation(2, 3)
	var target Location

	dist, err := location1.DistanceTo(target)
	assert.Error(t, err)
	assert.Equal(t, 0, dist)
}
