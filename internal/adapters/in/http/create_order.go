package http

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"delivery/internal/adapters/in/http/problems"
	"delivery/internal/core/application/usecases/commands"
	"delivery/internal/pkg/errs"
)

func (s *Server) CreateOrder(c echo.Context) error {
	cmd, err := commands.NewCreateOrderCommand(uuid.New(), "Несуществующая", 5)
	if err != nil {
		return problems.NewBadRequest(err.Error())
	}

	err = s.createOrderCommandHandler.Handle(c.Request().Context(), cmd)
	if err != nil {
		if errors.Is(err, errs.ErrObjectNotFound) {
			return problems.NewNotFound(err.Error())
		}
		return problems.NewConflict(err.Error(), "/")
	}

	return c.JSON(http.StatusOK, nil)
}
