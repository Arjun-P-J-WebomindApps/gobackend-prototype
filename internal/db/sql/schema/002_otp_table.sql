-- +goose Up
CREATE TABLE user_otps (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    opt_code VARCHAR NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    is_used BOOLEAN DEFAULT FALSE, -- if logged out and then try to reuse same otp before expiry will not allow
    created_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE user_opts;