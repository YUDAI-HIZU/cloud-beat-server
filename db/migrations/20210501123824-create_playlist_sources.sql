
-- +migrate Up
CREATE TABLE playlist_sources (
  playlist_id BIGINT(20) NOT NULL,
  track_id    BIGINT(20) NOT NULL,
  PRIMARY KEY (playlist_id, track_id),
  CONSTRAINT fk_playlist_sources_playlist_id FOREIGN KEY (playlist_id) REFERENCES playlists (id) ON DELETE CASCADE,
  CONSTRAINT fk_playlist_sources_track_id FOREIGN KEY (track_id) REFERENCES tracks (id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE playlist_sources;
