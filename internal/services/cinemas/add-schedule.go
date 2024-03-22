package cinemas

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/cinema-ticket/internal/helpers"
	"github.com/wildanfaz/cinema-ticket/internal/models"
)

func (s *ImplementServices) AddSchedule(c echo.Context) error {
	var (
		resp    = helpers.NewResponse()
		payload = new(models.Schedule)
	)

	err := c.Bind(payload)
	if err != nil {
		s.log.Errorln("Bind payload error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Unable to bind payload"))
	}

	err = validation.ValidateStruct(payload,
		validation.Field(&payload.MovieId, validation.Required, validation.Min(1)),
		validation.Field(&payload.ScheduleAt, validation.Required),
	)
	if err != nil {
		s.log.Errorln("Validate payload error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage(err.Error()))
	}

	err = s.cinemas.AddSchedule(c.Request().Context(), payload)
	if err != nil {
		s.log.Errorln("Add schedule error:", err)
		return c.JSON(http.StatusInternalServerError, resp.AsError().
			WithMessage("Unable to add schedule"))
	}

	return c.JSON(http.StatusOK, resp.WithMessage("Add schedule success"))
}
