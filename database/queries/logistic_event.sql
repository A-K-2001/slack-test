
-- name: GetLogisticEventsById :one
SELECT * FROM logistic_events WHERE id = $1;

-- name: GetLogisticEventsByRiderId :many
SELECT * FROM logistic_events WHERE rider_id = $1;