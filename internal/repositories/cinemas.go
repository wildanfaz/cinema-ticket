package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wildanfaz/cinema-ticket/internal/models"
)

type ImplementCinemas struct {
	db *pgxpool.Conn
}

type Cinemas interface {
	AddMovie(ctx context.Context, payload *models.Movie) error
	AddSchedule(ctx context.Context, payload *models.Schedule) error
	AddSeat(ctx context.Context, payload *models.Seat) error
	BookingTicket(ctx context.Context, payload *models.Ticket) error
	DeleteSchedule(ctx context.Context, id int) error
	UpdateSchedule(ctx context.Context, payload *models.Schedule) error
	ListSeats(ctx context.Context, scheduleId int) (models.Seats, error)
	ListMovies(ctx context.Context, payload *models.Movie) (models.ListMoviesResponse, error)
}

func NewCinemasRepository(db *pgxpool.Conn) Cinemas {
	return &ImplementCinemas{
		db: db,
	}
}

func (r *ImplementCinemas) AddMovie(ctx context.Context, payload *models.Movie) error {
	q := `
	INSERT INTO movies
	(title, description, price, duration)
	VALUES
	($1, $2, $3, $4)
	`

	_, err := r.db.Exec(ctx, q, payload.Title, payload.Description, payload.Price, payload.Duration)

	return err
}

func (r *ImplementCinemas) AddSchedule(ctx context.Context, payload *models.Schedule) error {
	q := `
	INSERT INTO schedules
	(movie_id, schedule_at)
	VALUES
	($1, $2)
	`

	_, err := r.db.Exec(ctx, q, payload.MovieId, payload.ScheduleAt)

	return err
}

func (r *ImplementCinemas) AddSeat(ctx context.Context, payload *models.Seat) error {
	q := `
	INSERT INTO seats
	(schedule_id, code)
	VALUES
	($1, $2)
	`

	_, err := r.db.Exec(ctx, q, payload.ScheduleId, payload.Code)

	return err
}

func (r *ImplementCinemas) BookingTicket(ctx context.Context, payload *models.Ticket) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return err
	}

	balance := 0
	err = tx.QueryRow(ctx, "SELECT balance FROM users WHERE id = $1 FOR UPDATE", payload.UserId).Scan(&balance)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	price, sold := 0, false
	tx.QueryRow(ctx, `
	SELECT m.price, se.sold FROM movies m
	JOIN schedules s ON s.movie_id = m.id
	JOIN seats se ON se.schedule_id = s.id
	WHERE se.id = $1 
	FOR UPDATE`, payload.SeatId).Scan(&price, &sold)

	if balance < price {
		tx.Rollback(ctx)
		return errors.New("Unable to booking ticket, please check user balance")
	}

	if sold {
		tx.Rollback(ctx)
		return errors.New("Unable to booking ticket, seat already sold")
	}

	_, err = tx.Exec(ctx, `
	UPDATE users
	SET balance = balance - $1
	WHERE id = $2
	`, price, payload.UserId)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, `
	UPDATE seats
	SET sold = true
	WHERE id = $1
	`, payload.SeatId)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}

func (r *ImplementCinemas) DeleteSchedule(ctx context.Context, id int) error {
	q1 := `
	SELECT count(0)
	FROM schedules
	WHERE id = $1
	`

	q2 := `
	DELETE FROM seats
	WHERE schedule_id = $1
	`

	q3 := `
	DELETE FROM schedules
	WHERE id = $1
	`

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	count := 0
	err = tx.QueryRow(ctx, q1, id).Scan(&count)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, q2, id)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, q3, id)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}

func (r *ImplementCinemas) UpdateSchedule(ctx context.Context, payload *models.Schedule) error {
	q := `
	UPDATE schedules
	SET schedule_at = $1
	WHERE id = $2
	`

	_, err := r.db.Exec(ctx, q, payload.ScheduleAt, payload.Id)

	return err
}

func (r *ImplementCinemas) ListSeats(ctx context.Context, scheduleId int) (models.Seats, error) {
	var seats models.Seats

	q := `
	SELECT id, schedule_id, code, sold, created_at, updated_at
	FROM seats
	WHERE schedule_id = $1
	`

	rows, err := r.db.Query(ctx, q, scheduleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var seat = new(models.Seat)
		err = rows.Scan(&seat.Id, &seat.ScheduleId, &seat.Code, &seat.Sold, &seat.CreatedAt, &seat.UpdatedAt)
		if err != nil {
			return nil, err
		}

		seat.ToLocal()
		seats = append(seats, *seat)
	}

	return seats, nil
}

func (r *ImplementCinemas) ListMovies(ctx context.Context, payload *models.Movie) (models.ListMoviesResponse, error) {
	var listMoviesResponse models.ListMoviesResponse

	q1 := `
	SELECT 
	id, title, description, price, duration, created_at, updated_at
	FROM movies
	`

	q2 := `
	SELECT
	id, movie_id, schedule_at, created_at, updated_at
	FROM schedules
	WHERE movie_id = $1
	`

	limit := 10
	offset := 0

	if payload.Pagination.PerPage > 0 {
		limit = payload.Pagination.PerPage
	}

	if payload.Pagination.Page > 1 {
		offset = (payload.Pagination.Page - 1) * limit
	}

	var values []any
	if payload.Title != "" {
		q1 = fmt.Sprintf("%s WHERE title @@ to_tsquery($1)", q1)
		values = append(values, payload.Title)
	}

	q1 = fmt.Sprintf("%s LIMIT %d OFFSET %d", q1, limit, offset)

	rows1, err := r.db.Query(ctx, q1, values...)
	if err != nil {
		return nil, err
	}
	defer rows1.Close()

	for rows1.Next() {
		var movie = new(models.Movie)
		err = rows1.Scan(
			&movie.Id, &movie.Title, &movie.Description,
			&movie.Price, &movie.Duration, &movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		movie.ToLocal()

		listMoviesResponse = append(listMoviesResponse, models.MovieWithSchedule{
			Movie:     *movie,
			Schedules: nil,
		})
	}

	for i, v := range listMoviesResponse {
		rows2, err := r.db.Query(ctx, q2, v.Movie.Id)
		if err != nil {
			return nil, err
		}
		defer rows2.Close()

		for rows2.Next() {
			var schedule = new(models.Schedule)
			err = rows2.Scan(
				&schedule.Id, &schedule.MovieId, &schedule.ScheduleAt,
				&schedule.CreatedAt, &schedule.UpdatedAt,
			)
			if err != nil {
				return nil, err
			}

			schedule.ToLocal()
			listMoviesResponse[i].Schedules = append(listMoviesResponse[i].Schedules, *schedule)
		}
	}

	return listMoviesResponse, nil
}
