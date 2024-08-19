-- name: GetTripByOrderId :one
SELECT * from logistic_trips WHERE order_id = $1;

-- name: GetTripByTripId :one
SELECT * from logistic_trips WHERE trip_id = $1;

-- name: GetTripsByPartner :many
SELECT * from logistic_trips WHERE partner = $1;

-- name: ListTripsByStatus :many
SELECT * FROM logistic_trips WHERE status = $1;