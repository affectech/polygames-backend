CREATE TABLE IF NOT EXISTS users
(
    id            bigserial primary key,
    username      text UNIQUE NOT NULL,
    email         text UNIQUE NOT NULL,
    password      text UNIQUE NOT NULL,
    gender        text,
    date_of_birth timestamptz,
    created_at    timestamptz NOT NULL DEFAULT now(),
    updated_at    timestamptz NOT NULL DEFAULT now(),
    role          int2 NOT NULL
);

CREATE TABLE IF NOT EXISTS genres
(
    id   bigserial PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE IF NOT EXISTS platforms
(
    id   bigserial PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE IF NOT EXISTS games
(
    id          bigserial PRIMARY KEY,
    user_id     int NOT NULL REFERENCES users (id),
    team_id     int NOT NULL,
    description text,
    genre_id    int NOT NULL REFERENCES genres (id),
    created_at  timestamptz NOT NULL DEFAULT now(),
    updated_at  timestamptz NOT NULL DEFAULT now(),
    likes       int,
    views       int,
    downloads   int,
    image       text NOT NULL,
    file        text NOT NULL
);

CREATE TABLE IF NOT EXISTS games_platforms
(
    game_id     int NOT NULL REFERENCES games (id),
    platform_id int NOT NULL REFERENCES platforms (id)
);
