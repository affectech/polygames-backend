CREATE TABLE IF NOT EXISTS games_platforms
(
    game_id     int NOT NULL REFERENCES games (id),
    platform_id int NOT NULL REFERENCES platforms (id)
);
