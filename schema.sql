CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  email text     NOT NULL UNIQUE,
  bio  text
);

CREATE TABLE books (
  id BIGSERIAL PRIMARY KEY,
  name text    NOT NULL
)
