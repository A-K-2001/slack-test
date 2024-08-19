-- name: GetRiders :many
SELECT * FROM rider;

-- name: GetRiderById :one
SELECT * FROM rider WHERE user_id = $1;

-- name: UpdateRiderLocation :one
UPDATE rider SET geohash = $1, battery = $2 WHERE user_id = $3 RETURNING *;

-- name: UpdateRiderStatus :one
UPDATE rider SET status = $2 WHERE user_id = $1 RETURNING *;

-- name: CreateRider :one
INSERT INTO rider (user_id, status, geohash, battery, vehicle_type) VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING RETURNING *;
