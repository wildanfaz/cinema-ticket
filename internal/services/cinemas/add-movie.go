package cinemas

import (
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/cinema-ticket/internal/helpers"
	"github.com/wildanfaz/cinema-ticket/internal/models"
)

func (s *ImplementServices) AddMovie(c echo.Context) error {
	var (
		resp    = helpers.NewResponse()
		payload = new(models.Movie)
	)

	err := c.Bind(payload)
	if err != nil {
		s.log.Errorln("Bind payload error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Unable to bind payload"))
	}

	err = validation.ValidateStruct(payload,
		validation.Field(&payload.Title, validation.Required, validation.Length(1, 255)),
		validation.Field(&payload.Description, validation.Required),
		validation.Field(&payload.Price, validation.Required, validation.Min(0)),
		validation.Field(&payload.Duration, validation.Required),
	)
	if err != nil {
		s.log.Errorln("Validate payload error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage(err.Error()))
	}

	_, err = time.ParseDuration(payload.Duration)
	if err != nil {
		s.log.Errorln("Parse duration error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Invalid duration"))
	}

	err = s.cinemas.AddMovie(c.Request().Context(), payload)
	if err != nil {
		s.log.Errorln("Add movie error:", err)
		return c.JSON(http.StatusInternalServerError, resp.AsError().
			WithMessage("Unable to add movie"))
	}

	return c.JSON(http.StatusOK, resp.WithMessage("Add movie success"))
}
