// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const insertSensorData = `-- name: InsertSensorData :one
INSERT INTO sensor_data (device_id, timestamp, temperature, humidity)
VALUES ($1, $2, $3, $4)
RETURNING id
`

type InsertSensorDataParams struct {
	DeviceID    pgtype.UUID
	Timestamp   pgtype.Timestamptz
	Temperature pgtype.Float8
	Humidity    pgtype.Float8
}

func (q *Queries) InsertSensorData(ctx context.Context, arg InsertSensorDataParams) (int64, error) {
	row := q.db.QueryRow(ctx, insertSensorData,
		arg.DeviceID,
		arg.Timestamp,
		arg.Temperature,
		arg.Humidity,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}
