package commands

type AssignOrderCommand struct{}

func NewAssignOrderCommand() (AssignOrderCommand, error) {
	return AssignOrderCommand{}, nil
}
