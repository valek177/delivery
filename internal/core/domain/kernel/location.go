package kernel

import (
	"math"
	"math/rand"

	"delivery/internal/pkg/errs"
)

const (
	minX = 1
	minY = 1
	maxX = 10
	maxY = 10
)

type Location struct {
	x int
	y int

	isSet bool
}

func NewLocation(x, y int) (Location, error) {
	if x < minX || x > maxX {
		return Location{}, errs.NewValueIsOutOfRangeError("x", x, minX, maxX)
	}

	if y < minY || x > maxY {
		return Location{}, errs.NewValueIsOutOfRangeError("y", y, minY, maxY)
	}

	return Location{
		x:     x,
		y:     y,
		isSet: true,
	}, nil
}

func NewRandomLocation() Location {
	x := minX + rand.Intn(maxX+1-minX)
	y := minY + rand.Intn(maxY+1-minY)

	return Location{
		x:     x,
		y:     y,
		isSet: true,
	}
}

func (l Location) X() int {
	return l.x
}

func (l Location) Y() int {
	return l.y
}

func (l Location) Equals(other Location) bool {
	return l == other
}

func (l Location) IsEmpty() bool {
	return !l.isSet
}

func (l Location) DistanceTo(target Location) (int, error) {
	if target.IsEmpty() {
		return 0, errs.NewValueIsRequiredError("location")
	}
	return int(
		math.Abs(float64(l.x-target.X())) +
			math.Abs(float64(l.y-target.Y())),
	), nil
}
