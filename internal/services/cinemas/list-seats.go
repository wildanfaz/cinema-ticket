package cinemas

import (
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/cinema-ticket/internal/helpers"
)

func (s *ImplementServices) ListSeats(c echo.Context) error {
	var (
		resp       = helpers.NewResponse()
		scheduleId = c.Param("schedule_id")
	)

	scheduleIdInt, err := strconv.Atoi(scheduleId)
	if err != nil {
		s.log.Errorln("Convert error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Invalid schedule id"))
	}

	err = validation.Validate(scheduleIdInt, validation.Required, validation.Min(1))
	if err != nil {
		s.log.Errorln("Validate payload error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage(err.Error()))
	}

	seats, err := s.cinemas.ListSeats(c.Request().Context(), scheduleIdInt)
	if err != nil {
		s.log.Errorln("List seats error:", err)
		return c.JSON(http.StatusInternalServerError, resp.AsError().
			WithMessage("Unable to list seats"))
	}

	return c.JSON(http.StatusOK, resp.WithMessage("Get list seats success").
		WithData(seats))
}
