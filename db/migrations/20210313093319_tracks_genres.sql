-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tracks_genres (
  track_id BIGINT(20) NOT NULL,
  genre_id BIGINT(20) NOT NULL,
  PRIMARY KEY (track_id, genre_id),
  CONSTRAINT fk_tracks_genres_track_id FOREIGN KEY (track_id) REFERENCES tracks (id),
  CONSTRAINT fk_tracks_genres_genre_id FOREIGN KEY (genre_id) REFERENCES genres (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE tracks_genres;
