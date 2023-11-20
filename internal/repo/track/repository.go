package track

import (
	"context"
	"database/sql"

	"github.com/Din4EE/soundsplayer/internal/repo"
	"github.com/Din4EE/soundsplayer/internal/repo/model"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repo.TrackRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]*model.Track, error) {
	builder := sq.Select("id", "title", "artist", "time", "cover", "audio").From("tracks").PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tracks []*model.Track
	for rows.Next() {
		track := &model.Track{}
		if err := rows.Scan(&track.Id, &track.Title, &track.Artist, &track.Time, &track.Cover, &track.Audio); err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}
