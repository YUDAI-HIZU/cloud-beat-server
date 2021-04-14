-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tracks (
  id              BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id         BIGINT(20) NOT NULL,
  title           VARCHAR(255) NOT NULL,
  thumbnail_path  VARCHAR(255) DEFAULT NULL,
  audio_path      VARCHAR(255) NOT NULL,
  description     TEXT NOT NULL,
  youtube_link    VARCHAR(255) NOT NULL,
  genre_id        BIGINT(20) NOT NULL,
  created_at      DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at      DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX(title),
  CONSTRAINT fk_tracks_user_id FOREIGN KEY (user_id) REFERENCES users (id),
  CONSTRAINT fk_tracks_genre_id FOREIGN KEY (genre_id) REFERENCES genres (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE tracks;
