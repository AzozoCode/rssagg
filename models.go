package main

import (
	"time"

	"github.com/azozocode/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	ApiKey   string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:       dbUser.ID,
		Name:     dbUser.Name,
		CreateAt: dbUser.CreateAt,
		UpdateAt: dbUser.UpdateAt,
		ApiKey:   dbUser.ApiKey,
	}
}
