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

type Feed struct {
	ID       uuid.UUID `json:"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	UserID   uuid.UUID `json:"user_id"`
	Name     string    `json:"name"`
	Url      string    `json:"url"`
}

type FeedFollow struct {
	ID       uuid.UUID `json:"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	UserID   uuid.UUID `json:"user_id"`
	FeedID   uuid.UUID `json:"feed_id"`
}

func databaseUserFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {

	return FeedFollow{
		ID:       dbFeedFollow.ID,
		CreateAt: dbFeedFollow.CreateAt,
		UpdateAt: dbFeedFollow.UpdateAt,
		UserID:   dbFeedFollow.UserID,
		FeedID:   dbFeedFollow.FeedID,
	}
}

func databaseUserFeedFollowToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {

	feed_follows := []FeedFollow{}
	for _, feed_follow := range dbFeedFollows {
		feed_follows = append(feed_follows, databaseUserFeedFollowToFeedFollow(feed_follow))

	}
	return feed_follows
}

func databaseUserFeedToFeed(dbFeed database.Feed) Feed {

	return Feed{
		ID:       dbFeed.ID,
		CreateAt: dbFeed.CreateAt,
		UpdateAt: dbFeed.UpdateAt,
		UserID:   dbFeed.UserID,
		Name:     dbFeed.Name,
		Url:      dbFeed.Url,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {

	feeds := []Feed{}
	for _, feed := range dbFeeds {
		feeds = append(feeds, databaseUserFeedToFeed(feed))
	}

	return feeds
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

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(dbPost database.Post) Post {

	var description *string

	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		CreateAt:    dbPost.CreateAt,
		UpdateAt:    dbPost.UpdateAt,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post {

	posts := []Post{}
	for _, post := range dbPosts {
		posts = append(posts, databasePostToPost(post))
	}

	return posts
}
