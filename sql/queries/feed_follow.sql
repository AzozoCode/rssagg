-- name: CreateFeedFollow :one

INSERT INTO feed_follow(id,create_at,update_at,user_id,feed_id)
VALUES($1,$2,$3,$4,$5)
RETURNING *;


-- name: GetFeedFollowByUserID :many
SELECT * FROM feed_follow WHERE user_id =$1;


-- name: DeleteFeedFollow :exec
DELETE FROM feed_follow WHERE id = $1 AND user_id = $2;