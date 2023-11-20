package track

import "github.com/Din4EE/soundsplayer/internal/repo"

type Service struct {
	trackRepo repo.TrackRepository
}

func NewService(trackRepo repo.TrackRepository) *Service {
	return &Service{
		trackRepo: trackRepo,
	}
}
