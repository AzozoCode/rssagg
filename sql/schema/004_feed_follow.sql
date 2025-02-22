-- +goose Up
    CREATE TABLE feed_follow(
    id UUID PRIMARY KEY NOT NULL,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE(user_id, feed_id)
    );





-- +goose Down

 DROP TABLE feed_follow;