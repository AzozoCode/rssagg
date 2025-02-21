-- +goose Up

CREATE TABLE users(
id UUID PRIMARY KEY NOT NULL,
create_at TIMESTAMP NOT NULL DEFAULT NOW(),
update_at TIMESTAMP NOT NULL DEFAULT NOW(),
name Text NOT NULL

);

-- +goose Down

DROP TABLE users;