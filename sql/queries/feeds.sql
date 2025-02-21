-- name: CreateFeed :one
INSERT INTO feeds(id,create_at,update_at,user_id,name,url)
VALUES($1,$2,$3,$4,$5,$6)
RETURNING *;


-- name: GetFeeds :many
SELECT * FROM feeds; 