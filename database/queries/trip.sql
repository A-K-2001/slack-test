-- name: InsertTrip :one
INSERT INTO trip (trip_type, dropoff_geohash, pickup_geohash, status, external_trip_id) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetRiderActiveTrips :many
SELECT  *
FROM trip
WHERE rider_id = $1
AND status != ALL(sqlc.arg('status')::trip_status[]) ORDER BY created_at DESC;

-- name: GetTripById :one
SELECT * FROM trip WHERE trip_id = $1;

-- name: GetRiderTripById :one
SELECT * FROM trip WHERE trip_id = $1 AND rider_id = $2;

-- name: GetRiderUnassignedTrip :one
SELECT * FROM trip WHERE trip_id = $1 AND rider_id IS NULL;

-- name: GetTripByExternalID :one
SELECT * FROM trip WHERE external_trip_id = $1;

-- name: GetUnassignedTripByExternalID :one
SELECT * FROM trip WHERE external_trip_id = $1 AND rider_id IS NULL;

-- name: UpdateTripRider :one
UPDATE trip SET rider_id=$2 WHERE trip_id = $1 RETURNING *;

-- name: UpdateTripPolyline :one
UPDATE trip SET encoded_polyline = $1 WHERE trip_id = $2 RETURNING *;

-- name: UpdateTripStatus :one
UPDATE trip SET status = $2 WHERE trip_id = $1 RETURNING *;

-- name: GetTripEvents :many
SELECT trip.trip_id, trip.pickup_geohash as trip_pickup_geohash, trip.dropoff_geohash as trip_dropoff_geohash, logistic_events.created_at, logistic_events.updated_at, event_type, event_battery_percentage, event_geohash, event_payload, checkpoint_order, checkpoint_type, checkpoint_geohash, arrived_event, visited_event FROM trip LEFT JOIN trip_checkpoint ON trip.trip_id = trip_checkpoint.trip_id LEFT JOIN logistic_events ON logistic_events.id = trip_checkpoint.arrived_event WHERE checkpoint_type != 'rerouted' AND trip.trip_id = $1 AND event_type IS NOT NULL ORDER BY checkpoint_order ASC;
