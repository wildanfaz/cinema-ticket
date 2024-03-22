CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS users (
    id bigserial NOT NULL,
    full_name varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    role varchar(255) NOT NULL DEFAULT 'user',
    balance bigint NOT NULL DEFAULT 0,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users (email);

DROP TRIGGER IF EXISTS users_updated_at_trigger ON users;

CREATE TRIGGER users_updated_at_trigger
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TABLE IF NOT EXISTS movies (
    id bigserial NOT NULL,
    title tsvector NOT NULL,
    description text,
    price int NOT NULL,
    duration varchar(50) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_movies_title ON movies USING gin(title);

DROP TRIGGER IF EXISTS movies_updated_at_trigger ON movies;

CREATE TRIGGER movies_updated_at_trigger
BEFORE UPDATE ON movies
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TABLE IF NOT EXISTS schedules (
    id bigserial NOT NULL,
    movie_id bigint NOT NULL,
    schedule_at timestamptz NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (movie_id) REFERENCES movies(id)
);

CREATE INDEX IF NOT EXISTS idx_schedules_movie_id ON schedules (movie_id);

DROP TRIGGER IF EXISTS schedules_updated_at_trigger ON schedules;

CREATE TRIGGER schedules_updated_at_trigger
BEFORE UPDATE ON schedules
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TABLE IF NOT EXISTS seats (
    id bigserial NOT NULL,
    schedule_id bigint NOT NULL,
    code varchar(5) NOT NULL,
    sold boolean NOT NULL DEFAULT false,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (schedule_id) REFERENCES schedules(id)
);

CREATE INDEX IF NOT EXISTS idx_seats_schedule_id ON seats (schedule_id);

DROP TRIGGER IF EXISTS seats_updated_at_trigger ON seats;

CREATE TRIGGER seats_updated_at_trigger
BEFORE UPDATE ON seats
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TABLE IF NOT EXISTS tickets (
    id bigserial NOT NULL,
    seat_id bigint NOT NULL,
    user_id bigint NOT NULL,
    order_by varchar(255) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (seat_id) REFERENCES seats(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_tickets_user_id ON tickets (user_id);

DROP TRIGGER IF EXISTS tickets_updated_at_trigger ON tickets;

CREATE TRIGGER tickets_updated_at_trigger
BEFORE UPDATE ON seats
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();