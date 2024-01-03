CREATE SCHEMA IF NOT EXISTS populate_cli;

CREATE TABLE IF NOT EXISTS populate_cli.movies (
    id INT PRIMARY KEY NOT NULL,
    title VARCHAR NOT NULL,
    release_date INT NOT NULL,
    genres VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);