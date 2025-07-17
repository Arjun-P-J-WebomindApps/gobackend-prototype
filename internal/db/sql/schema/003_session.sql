-- +goose Up
CREATE TABLE sessions (
    session_id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    ip_address TEXT,

    INDEX idx_sessions_user_id (user_id)
)

-- +goose Down
DROP TABLE sessions;