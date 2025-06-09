package kernel

import (
	"math"
	"math/rand"

	"delivery/internal/pkg/errs"
)

const (
	MIN_X = 1
	MIN_Y = 1
	MAX_X = 10
	MAX_Y = 10
)

type Location struct {
	x uint8
	y uint8

	isSet bool
}

func NewLocation(x, y uint8) (Location, error) {
	if x < MIN_X || x > MAX_X {
		return Location{}, errs.NewValueIsInvalidError("x")
	}

	if y < MIN_Y || x > MAX_Y {
		return Location{}, errs.NewValueIsInvalidError("y")
	}

	return Location{
		x:     x,
		y:     y,
		isSet: true,
	}, nil
}

func NewRandomLocation() Location {
	x := uint8(rand.Intn(MAX_X+1-MIN_X) + MIN_X)
	y := uint8(rand.Intn(MAX_Y+1-MIN_Y) + MIN_Y)

	return Location{
		x:     x,
		y:     y,
		isSet: true,
	}
}

func (l Location) X() uint8 {
	return l.x
}

func (l Location) Y() uint8 {
	return l.y
}

func (l Location) Equals(other Location) bool {
	return l == other
}

func (l Location) IsEmpty() bool {
}

func (l Location) IsValid() bool {
}

func (l Location) DistanceTo(target Location) (int, error) {
	if !target.IsValid() {
		return 0, errs.NewValueIsRequiredError("location")
	}
	return int(
		math.Abs(float64(l.x-target.X())) +
			math.Abs(float64(l.y-target.Y())),
	), nil
}
