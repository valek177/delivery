package http

import (
	"errors"
	"net/http"

	"delivery/internal/adapters/in/http/problems"
	"delivery/internal/core/application/usecases/commands"
	"delivery/internal/generated/servers"
	"delivery/internal/pkg/errs"

	"github.com/labstack/echo/v4"
)

func (s *Server) CreateCourier(c echo.Context) error {
	var courier servers.NewCourier
	if err := c.Bind(&courier); err != nil {
		return problems.NewBadRequest("invalid JSON body: " + err.Error())
	}

	createCourierCommand, err := commands.NewCreateCourierCommand(courier.Name, courier.Speed)
	if err != nil {
		return problems.NewBadRequest(err.Error())
	}

	err = s.createCourierCommandHandler.Handle(c.Request().Context(), createCourierCommand)
	if err != nil {
		if errors.Is(err, errs.ErrObjectNotFound) {
			return problems.NewNotFound(err.Error())
		}
		return problems.NewConflict(err.Error(), "/")
	}

	return c.JSON(http.StatusOK, nil)
}
