package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/cinema-ticket/internal/helpers"
	"github.com/wildanfaz/cinema-ticket/internal/pkg"
	"github.com/wildanfaz/cinema-ticket/internal/repositories"
)

func Auth(log *logrus.Logger, secretKey []byte, users repositories.Users, roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var (
				resp = helpers.NewResponse()
			)

			auth := c.Request().Header.Get("Authorization")
			if auth == "" {
				log.Errorln("Authorization header not found")
				return c.JSON(http.StatusUnauthorized, resp.AsError().
					WithMessage("Unauthorized"))
			}

			bearerToken := strings.Split(auth, " ")
			if len(bearerToken) < 2 {
				log.Errorln("Bearer token not found")
				return c.JSON(http.StatusUnauthorized, resp.AsError().
					WithMessage("Unauthorized"))
			}

			token := bearerToken[1]

			claims, err := pkg.ValidateToken(token, secretKey)
			if err != nil {
				log.Errorln("Validate token error:", err)
				return c.JSON(http.StatusUnauthorized, resp.AsError().
					WithMessage("Unauthorized"))
			}

			profile, err := users.Profile(c.Request().Context(), claims.Email)
			if err != nil {
				log.Errorln("Profile error:", err)
				return c.JSON(http.StatusUnauthorized, resp.AsError().
					WithMessage("Unauthorized"))
			}

			if profile == nil {
				log.Errorln("Profile not found")
				return c.JSON(http.StatusUnauthorized, resp.AsError().
					WithMessage("Unauthorized"))
			}

			roleValid := false
			for _, role := range roles {
				if profile.Role == role {
					roleValid = true
					break
				}
			}

			if !roleValid {
				log.Errorln("Role not valid")
				return c.JSON(http.StatusUnauthorized, resp.AsError().
					WithMessage("Unauthorized"))
			}

			c.Set("email", claims.Email)
			c.Set("user_id", profile.Id)

			return next(c)
		}
	}
}
