-- +goose Up
CREATE TABLE users (
    id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    PRIMARY KEY(id),
    api_key VARCHAR(64) NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE users;