package model

import (
	"database/sql"
	"time"
)

type Track struct {
	Id        uint64         `db:"id"`
	Audio     string         `db:"audio"`
	Title     sql.NullString `db:"title"`
	Artist    sql.NullString `db:"artist"`
	Time      uint64         `db:"time"`
	Cover     sql.NullString `db:"cover"`
	CreatedAt time.Time      `db:"created_at"`
}
