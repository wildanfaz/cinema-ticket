package users

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/cinema-ticket/configs"
	"github.com/wildanfaz/cinema-ticket/internal/repositories"
)

type ImplementServices struct {
	users  repositories.Users
	log    *logrus.Logger
	config *configs.Config
}

type Services interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}

func New(
	users repositories.Users,
	log *logrus.Logger,
	config *configs.Config,
) Services {
	return &ImplementServices{
		users:  users,
		log:    log,
		config: config,
	}
}
