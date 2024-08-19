-- name: FindAllTripCheckpoints :many
SELECT  *
FROM trip_checkpoint
WHERE trip_id = $1 AND checkpoint_type != 'rerouted'
ORDER BY checkpoint_order ASC;

-- name: InsertTripCheckpoint :one
INSERT INTO trip_checkpoint (trip_id, checkpoint_order, checkpoint_type, checkpoint_geohash, aggregated_trip_distance, aggregated_estimated_time, arrived_event, visited_event) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: UpdateTripCheckpointArrivedEvent :one
UPDATE trip_checkpoint SET arrived_event = $3 WHERE trip_id = $1 AND checkpoint_order <= $2 AND arrived_event IS NULL AND checkpoint_type != 'rerouted' RETURNING *;

-- name: UpdateTripCheckpointVisitedEvent :one
UPDATE trip_checkpoint SET visited_event = $3 WHERE trip_id = $1 AND checkpoint_order <= $2 AND visited_event IS NULL AND checkpoint_type != 'rerouted' RETURNING *;

-- name: MarkTripCheckpointsDetoured :many
UPDATE trip_checkpoint SET checkpoint_type = $3 WHERE trip_id = $1 AND checkpoint_order > $2 AND arrived_event IS NULL AND checkpoint_type != 'rerouted' RETURNING *;

-- name: GetNewCheckpointOrder :one
SELECT max(distinct checkpoint_order) + 1 as checkpoint_order  FROM trip_checkpoint WHERE trip_id=$1;

-- name: GetCheckpointByTypeForTripId :many
SELECT * FROM trip_checkpoint WHERE trip_id = $1 AND checkpoint_type = $2;
