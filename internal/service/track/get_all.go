package track

import (
	"context"
	"fmt"

	"github.com/Din4EE/soundsplayer/internal/service/model"
)

func (s *Service) GetAll(ctx context.Context) ([]*model.Track, error) {
	tracks, err := s.trackRepo.GetAll(ctx)

	srvTracks := make([]*model.Track, 0, len(tracks))

	for _, track := range tracks {
		srvTrack := &model.Track{
			Id:            track.Id,
			Audio:         track.Audio,
			Time:          track.Time,
			FormattedTime: fmt.Sprintf("%02d:%02d", track.Time/60, track.Time%60),
		}

		if track.Title.Valid {
			srvTrack.Title = track.Title.String
		} else {
			srvTrack.Title = "Unknown Title"
		}

		if track.Artist.Valid {
			srvTrack.Artist = track.Artist.String
		} else {
			srvTrack.Artist = "Unknown Artist"
		}

		if track.Cover.Valid {
			srvTrack.Cover = track.Cover.String
		} else {
			srvTrack.Cover = "assets/img/no_icon.png"
		}

		srvTracks = append(srvTracks, srvTrack)
	}

	if err != nil {
		return nil, err
	}
	return srvTracks, nil
}
