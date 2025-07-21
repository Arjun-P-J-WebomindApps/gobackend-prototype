-- +goose Up


-- Users---------------------------------------------------------------
CREATE TABLE users (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  username VARCHAR NOT NULL UNIQUE,
  email VARCHAR NOT NULL UNIQUE,
  password VARCHAR NOT NULL, -- salted and hashed
  mobile VARCHAR NOT NULL UNIQUE,
  role INT NOT NULL, -- 0=Admin, 1=Regular
  is_active BOOLEAN NOT NULL DEFAULT true,
  max_sessions INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);

-- OTPS---------------------------------------------------------------
CREATE TABLE user_otps (
  id SERIAL PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  otp_code VARCHAR NOT NULL,
  expires_at TIMESTAMP NOT NULL,
  is_used BOOLEAN NOT NULL DEFAULT FALSE, -- if logged out and then try to reuse same otp before expiry will not allow
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Allow only for one otp per user
CREATE INDEX idx_user_otps_user_id ON user_otps(user_id);

-- Session-------------------------------------------------------------
CREATE TABLE user_sessions (
  session_id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  ip_address TEXT NOT NULL,
  user_agent TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  expires_at TIMESTAMP NOT NULL,
  revoked_at TIMESTAMP  -- null means still active
);

CREATE INDEX idx_sessions_user_id on user_sessions(user_id);
CREATE INDEX idx_user_session_expires_at on user_sessions(expires_at);

--Refresh Token---------------------------------------------------------
CREATE TABLE refresh_tokens (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  session_id UUID NOT NULL REFERENCES user_sessions(session_id) ON DELETE CASCADE,
  token_hash TEXT NOT NULL UNIQUE, -- hashed version of refresh token
  ip_address TEXT,
  user_agent TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  expires_at TIMESTAMP NOT NULL,
  revoked_at TIMESTAMP,  -- null means still active
  replaced_by UUID REFERENCES refresh_tokens(id) -- for rotation
);

-- +goose Down
DROP TABLE refresh_tokens;
DROP TABLE user_sessions;
DROP TABLE user_otps;
DROP TABLE users;


