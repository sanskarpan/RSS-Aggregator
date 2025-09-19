-- +goose Up 
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    -- this will make sure that we can never have two instaneces of the same user-feed follow
    -- user can follow a certain instance of a feed once
    UNIQUE(user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;