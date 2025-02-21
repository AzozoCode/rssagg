-- +goose Up

CREATE TABLE feeds(
    id UUID PRIMARY KEY NOT NULL,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL
);


-- +goose Down

DROP TABLE feeds;



