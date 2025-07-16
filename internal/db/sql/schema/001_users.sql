-- +goose Up

CREATE TABLE users (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  username VARCHAR NOT NULL UNIQUE,
  email VARCHAR NOT NULL UNIQUE,
  password VARCHAR NOT NULL, -- salted and hashed
  mobile INT NOT NULL UNIQUE,
  role INT NOT NULL, -- 0=Admin, 1=Regular
  is_active BOOLEAN NOT NULL DEFAULT true,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);


-- +goose Down

DROP TABLE users;