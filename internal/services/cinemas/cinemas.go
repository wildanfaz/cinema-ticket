package cinemas

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/cinema-ticket/configs"
	"github.com/wildanfaz/cinema-ticket/internal/repositories"
)

type ImplementServices struct {
	cinemas repositories.Cinemas
	users   repositories.Users
	log     *logrus.Logger
	config  *configs.Config
}

type Services interface {
	AddMovie(c echo.Context) error
	AddSchedule(c echo.Context) error
	AddSeat(c echo.Context) error
	BookingTicket(c echo.Context) error
	DeleteSchedule(c echo.Context) error
	ListMovies(c echo.Context) error
	ListSeats(c echo.Context) error
	UpdateSchedule(c echo.Context) error
}

func New(
	cinemas repositories.Cinemas,
	users repositories.Users,
	log *logrus.Logger,
	config *configs.Config,
) Services {
	return &ImplementServices{
		cinemas: cinemas,
		users:   users,
		log:     log,
		config:  config,
	}
}
