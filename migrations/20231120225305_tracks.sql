-- +goose Up
-- +goose StatementBegin
CREATE TABLE tracks
(
    id         BIGSERIAL PRIMARY KEY,
    title      TEXT,
    artist     TEXT,
    audio      TEXT NOT NULL,
    cover      TEXT,
    time       INT,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

INSERT INTO tracks (id, title, artist, audio, cover, time) VALUES
(1, null, 'Crystal Castles', 'assets/audio/1106905450.mp3', null, 241),
(2, null, 'Crystal Castles', 'assets/audio/1826884532.mp3', null, 207),
(3, null, 'Crystal Castles', 'assets/audio/3572246108.mp3', null, 251);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tracks;
-- +goose StatementEnd
