// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: trip_checkpoint.sql

package models

import (
	"context"

	"github.com/google/uuid"
)

const findAllTripCheckpoints = `-- name: FindAllTripCheckpoints :many
SELECT  id, trip_id, checkpoint_order, checkpoint_type, arrived_event, visited_event, checkpoint_geohash, aggregated_trip_distance, aggregated_estimated_time, created_at, updated_at
FROM trip_checkpoint
WHERE trip_id = $1 AND checkpoint_type != 'rerouted'
ORDER BY checkpoint_order ASC
`

func (q *Queries) FindAllTripCheckpoints(ctx context.Context, tripID uuid.UUID) ([]TripCheckpoint, error) {
	rows, err := q.db.Query(ctx, findAllTripCheckpoints, tripID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TripCheckpoint{}
	for rows.Next() {
		var i TripCheckpoint
		if err := rows.Scan(
			&i.ID,
			&i.TripID,
			&i.CheckpointOrder,
			&i.CheckpointType,
			&i.ArrivedEvent,
			&i.VisitedEvent,
			&i.CheckpointGeohash,
			&i.AggregatedTripDistance,
			&i.AggregatedEstimatedTime,
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

const getCheckpointByTypeForTripId = `-- name: GetCheckpointByTypeForTripId :many
SELECT id, trip_id, checkpoint_order, checkpoint_type, arrived_event, visited_event, checkpoint_geohash, aggregated_trip_distance, aggregated_estimated_time, created_at, updated_at FROM trip_checkpoint WHERE trip_id = $1 AND checkpoint_type = $2
`

type GetCheckpointByTypeForTripIdParams struct {
	TripID         uuid.UUID          `db:"trip_id" json:"trip_id"`
	CheckpointType TripCheckpointType `db:"checkpoint_type" json:"checkpoint_type"`
}

func (q *Queries) GetCheckpointByTypeForTripId(ctx context.Context, arg GetCheckpointByTypeForTripIdParams) ([]TripCheckpoint, error) {
	rows, err := q.db.Query(ctx, getCheckpointByTypeForTripId, arg.TripID, arg.CheckpointType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TripCheckpoint{}
	for rows.Next() {
		var i TripCheckpoint
		if err := rows.Scan(
			&i.ID,
			&i.TripID,
			&i.CheckpointOrder,
			&i.CheckpointType,
			&i.ArrivedEvent,
			&i.VisitedEvent,
			&i.CheckpointGeohash,
			&i.AggregatedTripDistance,
			&i.AggregatedEstimatedTime,
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

const getNewCheckpointOrder = `-- name: GetNewCheckpointOrder :one
SELECT max(distinct checkpoint_order) + 1 as checkpoint_order  FROM trip_checkpoint WHERE trip_id=$1
`

func (q *Queries) GetNewCheckpointOrder(ctx context.Context, tripID uuid.UUID) (int32, error) {
	row := q.db.QueryRow(ctx, getNewCheckpointOrder, tripID)
	var checkpoint_order int32
	err := row.Scan(&checkpoint_order)
	return checkpoint_order, err
}

const insertTripCheckpoint = `-- name: InsertTripCheckpoint :one
INSERT INTO trip_checkpoint (trip_id, checkpoint_order, checkpoint_type, checkpoint_geohash, aggregated_trip_distance, aggregated_estimated_time, arrived_event, visited_event) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, trip_id, checkpoint_order, checkpoint_type, arrived_event, visited_event, checkpoint_geohash, aggregated_trip_distance, aggregated_estimated_time, created_at, updated_at
`

type InsertTripCheckpointParams struct {
	TripID                  uuid.UUID          `db:"trip_id" json:"trip_id"`
	CheckpointOrder         int32              `db:"checkpoint_order" json:"checkpoint_order"`
	CheckpointType          TripCheckpointType `db:"checkpoint_type" json:"checkpoint_type"`
	CheckpointGeohash       string             `db:"checkpoint_geohash" json:"checkpoint_geohash"`
	AggregatedTripDistance  int32              `db:"aggregated_trip_distance" json:"aggregated_trip_distance"`
	AggregatedEstimatedTime int32              `db:"aggregated_estimated_time" json:"aggregated_estimated_time"`
	ArrivedEvent            *uuid.UUID         `db:"arrived_event" json:"arrived_event"`
	VisitedEvent            *uuid.UUID         `db:"visited_event" json:"visited_event"`
}

func (q *Queries) InsertTripCheckpoint(ctx context.Context, arg InsertTripCheckpointParams) (TripCheckpoint, error) {
	row := q.db.QueryRow(ctx, insertTripCheckpoint,
		arg.TripID,
		arg.CheckpointOrder,
		arg.CheckpointType,
		arg.CheckpointGeohash,
		arg.AggregatedTripDistance,
		arg.AggregatedEstimatedTime,
		arg.ArrivedEvent,
		arg.VisitedEvent,
	)
	var i TripCheckpoint
	err := row.Scan(
		&i.ID,
		&i.TripID,
		&i.CheckpointOrder,
		&i.CheckpointType,
		&i.ArrivedEvent,
		&i.VisitedEvent,
		&i.CheckpointGeohash,
		&i.AggregatedTripDistance,
		&i.AggregatedEstimatedTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const markTripCheckpointsDetoured = `-- name: MarkTripCheckpointsDetoured :many
UPDATE trip_checkpoint SET checkpoint_type = $3 WHERE trip_id = $1 AND checkpoint_order > $2 AND arrived_event IS NULL AND checkpoint_type != 'rerouted' RETURNING id, trip_id, checkpoint_order, checkpoint_type, arrived_event, visited_event, checkpoint_geohash, aggregated_trip_distance, aggregated_estimated_time, created_at, updated_at
`

type MarkTripCheckpointsDetouredParams struct {
	TripID          uuid.UUID          `db:"trip_id" json:"trip_id"`
	CheckpointOrder int32              `db:"checkpoint_order" json:"checkpoint_order"`
	CheckpointType  TripCheckpointType `db:"checkpoint_type" json:"checkpoint_type"`
}

func (q *Queries) MarkTripCheckpointsDetoured(ctx context.Context, arg MarkTripCheckpointsDetouredParams) ([]TripCheckpoint, error) {
	rows, err := q.db.Query(ctx, markTripCheckpointsDetoured, arg.TripID, arg.CheckpointOrder, arg.CheckpointType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TripCheckpoint{}
	for rows.Next() {
		var i TripCheckpoint
		if err := rows.Scan(
			&i.ID,
			&i.TripID,
			&i.CheckpointOrder,
			&i.CheckpointType,
			&i.ArrivedEvent,
			&i.VisitedEvent,
			&i.CheckpointGeohash,
			&i.AggregatedTripDistance,
			&i.AggregatedEstimatedTime,
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

const updateTripCheckpointArrivedEvent = `-- name: UpdateTripCheckpointArrivedEvent :one
UPDATE trip_checkpoint SET arrived_event = $3 WHERE trip_id = $1 AND checkpoint_order <= $2 AND arrived_event IS NULL AND checkpoint_type != 'rerouted' RETURNING id, trip_id, checkpoint_order, checkpoint_type, arrived_event, visited_event, checkpoint_geohash, aggregated_trip_distance, aggregated_estimated_time, created_at, updated_at
`

type UpdateTripCheckpointArrivedEventParams struct {
	TripID          uuid.UUID  `db:"trip_id" json:"trip_id"`
	CheckpointOrder int32      `db:"checkpoint_order" json:"checkpoint_order"`
	ArrivedEvent    *uuid.UUID `db:"arrived_event" json:"arrived_event"`
}

func (q *Queries) UpdateTripCheckpointArrivedEvent(ctx context.Context, arg UpdateTripCheckpointArrivedEventParams) (TripCheckpoint, error) {
	row := q.db.QueryRow(ctx, updateTripCheckpointArrivedEvent, arg.TripID, arg.CheckpointOrder, arg.ArrivedEvent)
	var i TripCheckpoint
	err := row.Scan(
		&i.ID,
		&i.TripID,
		&i.CheckpointOrder,
		&i.CheckpointType,
		&i.ArrivedEvent,
		&i.VisitedEvent,
		&i.CheckpointGeohash,
		&i.AggregatedTripDistance,
		&i.AggregatedEstimatedTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTripCheckpointVisitedEvent = `-- name: UpdateTripCheckpointVisitedEvent :one
UPDATE trip_checkpoint SET visited_event = $3 WHERE trip_id = $1 AND checkpoint_order <= $2 AND visited_event IS NULL AND checkpoint_type != 'rerouted' RETURNING id, trip_id, checkpoint_order, checkpoint_type, arrived_event, visited_event, checkpoint_geohash, aggregated_trip_distance, aggregated_estimated_time, created_at, updated_at
`

type UpdateTripCheckpointVisitedEventParams struct {
	TripID          uuid.UUID  `db:"trip_id" json:"trip_id"`
	CheckpointOrder int32      `db:"checkpoint_order" json:"checkpoint_order"`
	VisitedEvent    *uuid.UUID `db:"visited_event" json:"visited_event"`
}

func (q *Queries) UpdateTripCheckpointVisitedEvent(ctx context.Context, arg UpdateTripCheckpointVisitedEventParams) (TripCheckpoint, error) {
	row := q.db.QueryRow(ctx, updateTripCheckpointVisitedEvent, arg.TripID, arg.CheckpointOrder, arg.VisitedEvent)
	var i TripCheckpoint
	err := row.Scan(
		&i.ID,
		&i.TripID,
		&i.CheckpointOrder,
		&i.CheckpointType,
		&i.ArrivedEvent,
		&i.VisitedEvent,
		&i.CheckpointGeohash,
		&i.AggregatedTripDistance,
		&i.AggregatedEstimatedTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
