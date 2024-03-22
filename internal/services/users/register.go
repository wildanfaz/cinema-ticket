package users

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/cinema-ticket/internal/helpers"
	"github.com/wildanfaz/cinema-ticket/internal/models"
	"github.com/wildanfaz/cinema-ticket/internal/pkg"
)

func (s *ImplementServices) Register(c echo.Context) error {
	var (
		resp    = helpers.NewResponse()
		payload = new(models.User)
	)

	err := c.Bind(payload)
	if err != nil {
		s.log.Errorln("Bind payload error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Unable to bind payload"))
	}

	err = validation.ValidateStruct(payload,
		validation.Field(&payload.Email, validation.Required, is.Email),
		validation.Field(&payload.Password, validation.Required, validation.Length(6, 50)),
		validation.Field(&payload.FullName, validation.Required, validation.Length(3, 50)),
	)
	if err != nil {
		s.log.Errorln("Validate payload error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage(err.Error()))
	}

	hashedPassword, err := pkg.HashPassword(payload.Password)
	if err != nil {
		s.log.Errorln("Hash password error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Unable to hash password"))
	}

	payload.Password = hashedPassword

	err = s.users.Register(c.Request().Context(), payload)
	if err != nil {
		s.log.Errorln("Register error:", err)
		return c.JSON(http.StatusInternalServerError, resp.AsError().
			WithMessage("Unable to register user"))
	}

	return c.JSON(http.StatusOK, resp.WithMessage("Register success"))
}
