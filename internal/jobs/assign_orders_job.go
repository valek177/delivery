package jobs

import (
	"context"

	"delivery/internal/core/application/usecases/commands"
	"delivery/internal/pkg/errs"

	"github.com/labstack/gommon/log"
	"github.com/robfig/cron/v3"
)

var _ cron.Job = &AssignOrdersJob{}

type AssignOrdersJob struct {
	assignOrdersCommandHandler commands.AssignOrderCommandHandler
}

func NewAssignOrdersJob(
	assignOrdersCommandHandler commands.AssignOrderCommandHandler,
) (cron.Job, error) {
	if assignOrdersCommandHandler == nil {
		return nil, errs.NewValueIsRequiredError("assignOrderCommandHandler")
	}

	return &AssignOrdersJob{
		assignOrdersCommandHandler: assignOrdersCommandHandler,
	}, nil
}

func (j *AssignOrdersJob) Run() {
	ctx := context.Background()
	command, err := commands.NewAssignOrderCommand()
	if err != nil {
		log.Error(err)
	}
	err = j.assignOrdersCommandHandler.Handle(ctx, command)
	if err != nil {
		log.Error(err)
	}
}
