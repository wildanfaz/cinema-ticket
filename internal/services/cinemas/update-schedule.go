package cinemas

import (
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/cinema-ticket/internal/helpers"
	"github.com/wildanfaz/cinema-ticket/internal/models"
)

func (s *ImplementServices) UpdateSchedule(c echo.Context) error {
	var (
		resp    = helpers.NewResponse()
		payload = new(models.Schedule)
		id      = c.Param("id")
	)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		s.log.Errorln("Convert error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Invalid id"))
	}

	err = c.Bind(payload)
	if err != nil {
		s.log.Errorln("Bind payload error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Unable to bind payload"))
	}

	payload.Id = idInt

	err = validation.ValidateStruct(payload,
		validation.Field(&payload.Id, validation.Required, validation.Min(1)),
		validation.Field(&payload.ScheduleAt, validation.Required),
	)
	if err != nil {
		s.log.Errorln("Validate payload error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage(err.Error()))
	}

	err = s.cinemas.UpdateSchedule(c.Request().Context(), payload)
	if err != nil {
		s.log.Errorln("Update schedule error:", err)
		return c.JSON(http.StatusInternalServerError, resp.AsError().
			WithMessage("Unable to update schedule"))
	}

	return c.JSON(http.StatusOK, resp.WithMessage("Update schedule success"))
}
