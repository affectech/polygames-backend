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


