package cinemas

import (
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/cinema-ticket/internal/helpers"
)

func (s *ImplementServices) DeleteSchedule(c echo.Context) error {
	var (
		resp = helpers.NewResponse()
		id   = c.Param("id")
	)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		s.log.Errorln("Convert error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Invalid id"))
	}

	err = validation.Validate(idInt, validation.Required, validation.Min(1))
	if err != nil {
		s.log.Errorln("Validate payload error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage(err.Error()))
	}

	err = s.cinemas.DeleteSchedule(c.Request().Context(), idInt)
	if err == pgx.ErrNoRows {
		s.log.Errorln("Delete schedule error:", err)
		return c.JSON(http.StatusNotFound, resp.AsError().
			WithMessage("Schedule not found"))
	}

	if err != nil {
		s.log.Errorln("Delete schedule error:", err)
		return c.JSON(http.StatusInternalServerError, resp.AsError().
			WithMessage("Unable to delete schedule"))
	}

	return c.JSON(http.StatusOK, resp.WithMessage("Delete schedule success"))
}
