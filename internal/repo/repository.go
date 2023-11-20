package repo

import (
	"context"

	"github.com/Din4EE/soundsplayer/internal/repo/model"
)

type TrackRepository interface {
	GetAll(ctx context.Context) ([]*model.Track, error)
}
