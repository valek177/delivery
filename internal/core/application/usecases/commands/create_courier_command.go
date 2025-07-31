package commands

import "delivery/internal/pkg/errs"

type CreateCourierCommand struct {
	name  string
	speed int
}

func (c CreateCourierCommand) Name() string {
	return c.name
}

func (c CreateCourierCommand) Speed() int {
	return c.speed
}

func NewCreateCourierCommand(name string, speed int) (CreateCourierCommand, error) {
	if name == "" {
		return CreateCourierCommand{}, errs.NewValueIsRequiredError("name")
	}
	if speed <= 0 {
		return CreateCourierCommand{}, errs.NewValueIsRequiredError("speed")
	}

	return CreateCourierCommand{
		name:  name,
		speed: speed,
	}, nil
}
