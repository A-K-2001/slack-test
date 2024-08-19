-- name: GetAppUpdates :many
SELECT * from app_updates;

-- name: GetAppUpdatesByPlatform :one
SELECT * from app_updates WHERE platform = $1;