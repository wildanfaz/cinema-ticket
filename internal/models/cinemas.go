package models

import "time"

type (
	Pagination struct {
		PerPage int `query:"per_page" json:"-"`
		Page    int `query:"page" json:"-"`
	}

	Movie struct {
		Id          int       `json:"id"`
		Title       string    `json:"title" query:"title"`
		Description string    `json:"description"`
		Price       int       `json:"price"`
		Duration    string    `json:"duration"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Pagination
	}

	Schedule struct {
		Id         int       `json:"id"`
		MovieId    int       `json:"movie_id"`
		ScheduleAt time.Time `json:"schedule_at"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	Schedules []Schedule

	MovieWithSchedule struct {
		Movie     Movie     `json:"movie"`
		Schedules Schedules `json:"schedules"`
	}

	ListMoviesResponse []MovieWithSchedule

	Seat struct {
		Id         int       `json:"id"`
		ScheduleId int       `json:"schedule_id"`
		Code       string    `json:"code"`
		Sold       bool      `json:"sold"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	Seats []Seat

	Ticket struct {
		Id        int       `json:"id"`
		SeatId    int       `json:"seat_id"`
		UserId    int       `json:"user_id"`
		OrderBy   string    `json:"order_by"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func (m *Movie) ToLocal() {
	m.CreatedAt = m.CreatedAt.Local()
	m.UpdatedAt = m.UpdatedAt.Local()
}

func (m *Schedule) ToLocal() {
	m.ScheduleAt = m.ScheduleAt.Local()
	m.CreatedAt = m.CreatedAt.Local()
	m.UpdatedAt = m.UpdatedAt.Local()
}

func (m *Seat) ToLocal() {
	m.CreatedAt = m.CreatedAt.Local()
	m.UpdatedAt = m.UpdatedAt.Local()
}

func (m *Ticket) ToLocal() {
	m.CreatedAt = m.CreatedAt.Local()
	m.UpdatedAt = m.UpdatedAt.Local()
}
