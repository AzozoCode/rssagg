// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(id,create_at,update_at,name,api_key) 
VALUES($1,$2,$3,$4,encode(sha256(random()::text::bytea),'hex'))
RETURNING id, create_at, update_at, name, api_key
`

type CreateUserParams struct {
	ID       uuid.UUID
	CreateAt time.Time
	UpdateAt time.Time
	Name     string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreateAt,
		arg.UpdateAt,
		arg.Name,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreateAt,
		&i.UpdateAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}

const getUserByAPIKey = `-- name: GetUserByAPIKey :one
SELECT id, create_at, update_at, name, api_key FROM users WHERE api_key = $1
`

func (q *Queries) GetUserByAPIKey(ctx context.Context, apiKey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByAPIKey, apiKey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreateAt,
		&i.UpdateAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}
