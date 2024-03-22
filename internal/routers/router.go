package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/cinema-ticket/configs"
	"github.com/wildanfaz/cinema-ticket/internal/middlewares"
	"github.com/wildanfaz/cinema-ticket/internal/pkg"
	"github.com/wildanfaz/cinema-ticket/internal/repositories"
	"github.com/wildanfaz/cinema-ticket/internal/services/cinemas"
	"github.com/wildanfaz/cinema-ticket/internal/services/health"
	"github.com/wildanfaz/cinema-ticket/internal/services/users"
)

func InitEchoRouter() {
	// configs
	config := configs.InitConfig()
	db := configs.InitPostgreSQL(config.DatabaseURL)

	// pkg
	log := pkg.InitLogger()

	// repositories
	usersRepo := repositories.NewUsersRepository(db)
	cinemasRepo := repositories.NewCinemasRepository(db)

	// services
	usersServices := users.New(usersRepo, log, config)
	cinemasServices := cinemas.New(cinemasRepo, usersRepo, log, config)

	e := echo.New()

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/health", health.HealthCheck)

	// users
	usersRoute := apiV1.Group("/users")

	usersRoute.POST("/register", usersServices.Register)
	usersRoute.POST("/login", usersServices.Login)

	// cinemas
	cinemasRoute := apiV1.Group("/cinemas")
	cinemasRoute.POST("/add-movie", cinemasServices.AddMovie, middlewares.Auth(log, config.JWTSecretKey, usersRepo, "admin"))
	cinemasRoute.POST("/add-schedule", cinemasServices.AddSchedule, middlewares.Auth(log, config.JWTSecretKey, usersRepo, "admin"))
	cinemasRoute.POST("/add-seat", cinemasServices.AddSeat, middlewares.Auth(log, config.JWTSecretKey, usersRepo, "admin"))
	cinemasRoute.POST("/booking-ticket", cinemasServices.BookingTicket, middlewares.Auth(log, config.JWTSecretKey, usersRepo, "user", "admin"))

	cinemasRoute.GET("/list-movies", cinemasServices.ListMovies, middlewares.Auth(log, config.JWTSecretKey, usersRepo, "user", "admin"))
	cinemasRoute.GET("/list-seats/:schedule_id", cinemasServices.ListSeats, middlewares.Auth(log, config.JWTSecretKey, usersRepo, "user", "admin"))

	cinemasRoute.PUT("/update-schedule/:id", cinemasServices.UpdateSchedule, middlewares.Auth(log, config.JWTSecretKey, usersRepo, "admin"))

	cinemasRoute.DELETE("/delete-schedule/:id", cinemasServices.DeleteSchedule, middlewares.Auth(log, config.JWTSecretKey, usersRepo, "admin"))

	e.Logger.Fatal(e.Start(config.AppPort))
}
