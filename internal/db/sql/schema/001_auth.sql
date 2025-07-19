-- +goose Up

CREATE TABLE users (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  username VARCHAR NOT NULL UNIQUE,
  email VARCHAR NOT NULL UNIQUE,
  password VARCHAR NOT NULL, -- salted and hashed
  mobile VARCHAR NOT NULL UNIQUE,
  role INT NOT NULL, -- 0=Admin, 1=Regular
  is_active BOOLEAN NOT NULL DEFAULT true,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);


CREATE TABLE user_otps (
  id SERIAL PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  otp_code VARCHAR NOT NULL,
  expires_at TIMESTAMP NOT NULL,
  is_used BOOLEAN NOT NULL DEFAULT FALSE, -- if logged out and then try to reuse same otp before expiry will not allow
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE user_sessions (
  session_id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP DEFAULT NOW(),
  expires_at TIMESTAMP NOT NULL,
  ip_address TEXT NOT NULL
);

CREATE INDEX idx_sessions_user_id on user_sessions(user_id);

-- +goose Down
DROP TABLE user_sessions;
DROP TABLE user_otps;
DROP TABLE users;


