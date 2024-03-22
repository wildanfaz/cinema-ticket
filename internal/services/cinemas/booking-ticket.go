package cinemas

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/cinema-ticket/internal/helpers"
	"github.com/wildanfaz/cinema-ticket/internal/models"
)

func (s *ImplementServices) BookingTicket(c echo.Context) error {
	var (
		resp    = helpers.NewResponse()
		payload = new(models.Ticket)
		userId  = c.Get("user_id").(int)
	)

	err := c.Bind(payload)
	if err != nil {
		s.log.Errorln("Bind payload error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Unable to bind payload"))
	}

	err = validation.ValidateStruct(payload,
		validation.Field(&payload.SeatId, validation.Required, validation.Min(1)),
		validation.Field(&payload.OrderBy, validation.Required),
	)
	if err != nil {
		s.log.Errorln("Validate payload error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage(err.Error()))
	}

	payload.UserId = userId

	err = s.cinemas.BookingTicket(c.Request().Context(), payload)
	if err != nil {
		s.log.Errorln("Booking ticket error:", err)
		return c.JSON(http.StatusInternalServerError, resp.AsError().
			WithMessage("Unable to booking ticket"))
	}

	return c.JSON(http.StatusOK, resp.WithMessage("Booking ticket success"))
}
