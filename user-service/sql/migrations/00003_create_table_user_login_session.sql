-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_login_session (
    id SERIAL,
    user_id BIGINT NOT NULL,
    public_key TEXT NOT NULL,
    rf_token TEXT NOT NULL,
    rf_token_used TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_login_session;
-- +goose StatementEnd
