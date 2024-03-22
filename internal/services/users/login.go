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

func (s *ImplementServices) Login(c echo.Context) error {
	var (
		resp    = helpers.NewResponse()
		payload = new(models.Login)
	)

	err := c.Bind(payload)
	if err != nil {
		s.log.Errorln("Bind payload error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Unable to bind payload"))
	}

	err = validation.ValidateStruct(payload,
		validation.Field(&payload.Email, validation.Required, is.Email),
		validation.Field(&payload.Password, validation.Required),
		validation.Field(&payload.ConfirmPassword, validation.Required),
	)
	if err != nil {
		s.log.Errorln("Validate payload error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage(err.Error()))
	}

	if payload.Password != payload.ConfirmPassword {
		s.log.Errorln("Password and confirm password not match")
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Password and confirm password not match"))
	}

	profile, err := s.users.Profile(c.Request().Context(), payload.Email)
	if err != nil {
		s.log.Errorln("Profile error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Invalid email or password"))
	}

	if profile == nil {
		s.log.Errorln("Profile not found")
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Invalid email or password"))
	}

	err = pkg.ComparePassword(payload.Password, profile.Password)
	if err != nil {
		s.log.Errorln("Compare password error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Invalid email or password"))
	}

	token, err := pkg.GenerateToken(&pkg.NewClaims{
		Email: payload.Email,
	}, s.config.JWTSecretKey)
	if err != nil {
		s.log.Errorln("Generate token error:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Unable to login"))
	}

	return c.JSON(http.StatusOK, resp.WithMessage("Login success").
		WithData(map[string]string{
			"token": token,
		}))
}
