// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: trip.sql

package models

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const getRiderActiveTrips = `-- name: GetRiderActiveTrips :many
SELECT  trip_id, external_trip_id, encoded_polyline, rider_id, trip_type, dropoff_geohash, pickup_geohash, status, created_at, updated_at
FROM trip
WHERE rider_id = $1
AND status != ALL($2::trip_status[]) ORDER BY created_at DESC
`

type GetRiderActiveTripsParams struct {
	RiderID *string      `db:"rider_id" json:"rider_id"`
	Status  []TripStatus `db:"status" json:"status"`
}

func (q *Queries) GetRiderActiveTrips(ctx context.Context, arg GetRiderActiveTripsParams) ([]Trip, error) {
	rows, err := q.db.Query(ctx, getRiderActiveTrips, arg.RiderID, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Trip{}
	for rows.Next() {
		var i Trip
		if err := rows.Scan(
			&i.TripID,
			&i.ExternalTripID,
			&i.EncodedPolyline,
			&i.RiderID,
			&i.TripType,
			&i.DropoffGeohash,
			&i.PickupGeohash,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRiderTripById = `-- name: GetRiderTripById :one
SELECT trip_id, external_trip_id, encoded_polyline, rider_id, trip_type, dropoff_geohash, pickup_geohash, status, created_at, updated_at FROM trip WHERE trip_id = $1 AND rider_id = $2
`

type GetRiderTripByIdParams struct {
	TripID  uuid.UUID `db:"trip_id" json:"trip_id"`
	RiderID *string   `db:"rider_id" json:"rider_id"`
}

func (q *Queries) GetRiderTripById(ctx context.Context, arg GetRiderTripByIdParams) (Trip, error) {
	row := q.db.QueryRow(ctx, getRiderTripById, arg.TripID, arg.RiderID)
	var i Trip
	err := row.Scan(
		&i.TripID,
		&i.ExternalTripID,
		&i.EncodedPolyline,
		&i.RiderID,
		&i.TripType,
		&i.DropoffGeohash,
		&i.PickupGeohash,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRiderUnassignedTrip = `-- name: GetRiderUnassignedTrip :one
SELECT trip_id, external_trip_id, encoded_polyline, rider_id, trip_type, dropoff_geohash, pickup_geohash, status, created_at, updated_at FROM trip WHERE trip_id = $1 AND rider_id IS NULL
`

func (q *Queries) GetRiderUnassignedTrip(ctx context.Context, tripID uuid.UUID) (Trip, error) {
	row := q.db.QueryRow(ctx, getRiderUnassignedTrip, tripID)
	var i Trip
	err := row.Scan(
		&i.TripID,
		&i.ExternalTripID,
		&i.EncodedPolyline,
		&i.RiderID,
		&i.TripType,
		&i.DropoffGeohash,
		&i.PickupGeohash,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTripByExternalID = `-- name: GetTripByExternalID :one
SELECT trip_id, external_trip_id, encoded_polyline, rider_id, trip_type, dropoff_geohash, pickup_geohash, status, created_at, updated_at FROM trip WHERE external_trip_id = $1
`

func (q *Queries) GetTripByExternalID(ctx context.Context, externalTripID *string) (Trip, error) {
	row := q.db.QueryRow(ctx, getTripByExternalID, externalTripID)
	var i Trip
	err := row.Scan(
		&i.TripID,
		&i.ExternalTripID,
		&i.EncodedPolyline,
		&i.RiderID,
		&i.TripType,
		&i.DropoffGeohash,
		&i.PickupGeohash,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTripById = `-- name: GetTripById :one
SELECT trip_id, external_trip_id, encoded_polyline, rider_id, trip_type, dropoff_geohash, pickup_geohash, status, created_at, updated_at FROM trip WHERE trip_id = $1
`

func (q *Queries) GetTripById(ctx context.Context, tripID uuid.UUID) (Trip, error) {
	row := q.db.QueryRow(ctx, getTripById, tripID)
	var i Trip
	err := row.Scan(
		&i.TripID,
		&i.ExternalTripID,
		&i.EncodedPolyline,
		&i.RiderID,
		&i.TripType,
		&i.DropoffGeohash,
		&i.PickupGeohash,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTripEvents = `-- name: GetTripEvents :many
SELECT trip.trip_id, trip.pickup_geohash as trip_pickup_geohash, trip.dropoff_geohash as trip_dropoff_geohash, logistic_events.created_at, logistic_events.updated_at, event_type, event_battery_percentage, event_geohash, event_payload, checkpoint_order, checkpoint_type, checkpoint_geohash, arrived_event, visited_event FROM trip LEFT JOIN trip_checkpoint ON trip.trip_id = trip_checkpoint.trip_id LEFT JOIN logistic_events ON logistic_events.id = trip_checkpoint.arrived_event WHERE checkpoint_type != 'rerouted' AND trip.trip_id = $1 AND event_type IS NOT NULL ORDER BY checkpoint_order ASC
`

type GetTripEventsRow struct {
	TripID                 uuid.UUID              `db:"trip_id" json:"trip_id"`
	TripPickupGeohash      string                 `db:"trip_pickup_geohash" json:"trip_pickup_geohash"`
	TripDropoffGeohash     string                 `db:"trip_dropoff_geohash" json:"trip_dropoff_geohash"`
	CreatedAt              pgtype.Timestamptz     `db:"created_at" json:"created_at"`
	UpdatedAt              pgtype.Timestamptz     `db:"updated_at" json:"updated_at"`
	EventType              NullLogisticEventType  `db:"event_type" json:"event_type"`
	EventBatteryPercentage *int32                 `db:"event_battery_percentage" json:"event_battery_percentage"`
	EventGeohash           *string                `db:"event_geohash" json:"event_geohash"`
	EventPayload           []byte                 `db:"event_payload" json:"event_payload"`
	CheckpointOrder        *int32                 `db:"checkpoint_order" json:"checkpoint_order"`
	CheckpointType         NullTripCheckpointType `db:"checkpoint_type" json:"checkpoint_type"`
	CheckpointGeohash      *string                `db:"checkpoint_geohash" json:"checkpoint_geohash"`
	ArrivedEvent           *uuid.UUID             `db:"arrived_event" json:"arrived_event"`
	VisitedEvent           *uuid.UUID             `db:"visited_event" json:"visited_event"`
}

func (q *Queries) GetTripEvents(ctx context.Context, tripID uuid.UUID) ([]GetTripEventsRow, error) {
	rows, err := q.db.Query(ctx, getTripEvents, tripID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetTripEventsRow{}
	for rows.Next() {
		var i GetTripEventsRow
		if err := rows.Scan(
			&i.TripID,
			&i.TripPickupGeohash,
			&i.TripDropoffGeohash,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.EventType,
			&i.EventBatteryPercentage,
			&i.EventGeohash,
			&i.EventPayload,
			&i.CheckpointOrder,
			&i.CheckpointType,
			&i.CheckpointGeohash,
			&i.ArrivedEvent,
			&i.VisitedEvent,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnassignedTripByExternalID = `-- name: GetUnassignedTripByExternalID :one
SELECT trip_id, external_trip_id, encoded_polyline, rider_id, trip_type, dropoff_geohash, pickup_geohash, status, created_at, updated_at FROM trip WHERE external_trip_id = $1 AND rider_id IS NULL
`

func (q *Queries) GetUnassignedTripByExternalID(ctx context.Context, externalTripID *string) (Trip, error) {
	row := q.db.QueryRow(ctx, getUnassignedTripByExternalID, externalTripID)
	var i Trip
	err := row.Scan(
		&i.TripID,
		&i.ExternalTripID,
		&i.EncodedPolyline,
		&i.RiderID,
		&i.TripType,
		&i.DropoffGeohash,
		&i.PickupGeohash,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertTrip = `-- name: InsertTrip :one
INSERT INTO trip (trip_type, dropoff_geohash, pickup_geohash, status, external_trip_id) VALUES ($1, $2, $3, $4, $5) RETURNING trip_id, external_trip_id, encoded_polyline, rider_id, trip_type, dropoff_geohash, pickup_geohash, status, created_at, updated_at
`

type InsertTripParams struct {
	TripType       TripType   `db:"trip_type" json:"trip_type"`
	DropoffGeohash string     `db:"dropoff_geohash" json:"dropoff_geohash"`
	PickupGeohash  string     `db:"pickup_geohash" json:"pickup_geohash"`
	Status         TripStatus `db:"status" json:"status"`
	ExternalTripID *string    `db:"external_trip_id" json:"external_trip_id"`
}

func (q *Queries) InsertTrip(ctx context.Context, arg InsertTripParams) (Trip, error) {
	row := q.db.QueryRow(ctx, insertTrip,
		arg.TripType,
		arg.DropoffGeohash,
		arg.PickupGeohash,
		arg.Status,
		arg.ExternalTripID,
	)
	var i Trip
	err := row.Scan(
		&i.TripID,
		&i.ExternalTripID,
		&i.EncodedPolyline,
		&i.RiderID,
		&i.TripType,
		&i.DropoffGeohash,
		&i.PickupGeohash,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTripPolyline = `-- name: UpdateTripPolyline :one
UPDATE trip SET encoded_polyline = $1 WHERE trip_id = $2 RETURNING trip_id, external_trip_id, encoded_polyline, rider_id, trip_type, dropoff_geohash, pickup_geohash, status, created_at, updated_at
`

type UpdateTripPolylineParams struct {
	EncodedPolyline *string   `db:"encoded_polyline" json:"encoded_polyline"`
	TripID          uuid.UUID `db:"trip_id" json:"trip_id"`
}

func (q *Queries) UpdateTripPolyline(ctx context.Context, arg UpdateTripPolylineParams) (Trip, error) {
	row := q.db.QueryRow(ctx, updateTripPolyline, arg.EncodedPolyline, arg.TripID)
	var i Trip
	err := row.Scan(
		&i.TripID,
		&i.ExternalTripID,
		&i.EncodedPolyline,
		&i.RiderID,
		&i.TripType,
		&i.DropoffGeohash,
		&i.PickupGeohash,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTripRider = `-- name: UpdateTripRider :one
UPDATE trip SET rider_id=$2 WHERE trip_id = $1 RETURNING trip_id, external_trip_id, encoded_polyline, rider_id, trip_type, dropoff_geohash, pickup_geohash, status, created_at, updated_at
`

type UpdateTripRiderParams struct {
	TripID  uuid.UUID `db:"trip_id" json:"trip_id"`
	RiderID *string   `db:"rider_id" json:"rider_id"`
}

func (q *Queries) UpdateTripRider(ctx context.Context, arg UpdateTripRiderParams) (Trip, error) {
	row := q.db.QueryRow(ctx, updateTripRider, arg.TripID, arg.RiderID)
	var i Trip
	err := row.Scan(
		&i.TripID,
		&i.ExternalTripID,
		&i.EncodedPolyline,
		&i.RiderID,
		&i.TripType,
		&i.DropoffGeohash,
		&i.PickupGeohash,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTripStatus = `-- name: UpdateTripStatus :one
UPDATE trip SET status = $2 WHERE trip_id = $1 RETURNING trip_id, external_trip_id, encoded_polyline, rider_id, trip_type, dropoff_geohash, pickup_geohash, status, created_at, updated_at
`

type UpdateTripStatusParams struct {
	TripID uuid.UUID  `db:"trip_id" json:"trip_id"`
	Status TripStatus `db:"status" json:"status"`
}

func (q *Queries) UpdateTripStatus(ctx context.Context, arg UpdateTripStatusParams) (Trip, error) {
	row := q.db.QueryRow(ctx, updateTripStatus, arg.TripID, arg.Status)
	var i Trip
	err := row.Scan(
		&i.TripID,
		&i.ExternalTripID,
		&i.EncodedPolyline,
		&i.RiderID,
		&i.TripType,
		&i.DropoffGeohash,
		&i.PickupGeohash,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
