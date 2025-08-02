package jobs

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/robfig/cron/v3"

	"delivery/internal/core/application/usecases/commands"
	"delivery/internal/pkg/errs"
)

var _ cron.Job = &MoveCouriersJob{}

type MoveCouriersJob struct {
	moveCouriersCommandHandler commands.MoveCouriersCommandHandler
}

func NewMoveCouriersJob(
	moveCouriersCommandHandler commands.MoveCouriersCommandHandler,
) (cron.Job, error) {
	if moveCouriersCommandHandler == nil {
		return nil, errs.NewValueIsRequiredError("moveCouriersCommandHandler")
	}

	return &MoveCouriersJob{
		moveCouriersCommandHandler: moveCouriersCommandHandler,
	}, nil
}

func (j *MoveCouriersJob) Run() {
	ctx := context.Background()
	command, err := commands.NewMoveCouriersCommand()
	if err != nil {
		log.Error(err)
	}
	err = j.moveCouriersCommandHandler.Handle(ctx, command)
	if err != nil {
		log.Error(err)
	}
}
