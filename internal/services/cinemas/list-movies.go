package cinemas

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/cinema-ticket/internal/helpers"
	"github.com/wildanfaz/cinema-ticket/internal/models"
)

func (s *ImplementServices) ListMovies(c echo.Context) error {
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

	movies, err := s.cinemas.ListMovies(c.Request().Context(), payload)
	if err != nil {
		s.log.Errorln("List movies error:", err)
		return c.JSON(http.StatusInternalServerError, resp.AsError().
			WithMessage("Unable to get list movies"))
	}

	return c.JSON(http.StatusOK, resp.WithMessage("Get list movies success").
		WithData(movies))
}
