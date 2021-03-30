
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE music_videos (
  id          BIGINT(20) NOT NULL PRIMARY KEY,
  user_id     BIGINT(20) NOT NULL,
  text        TEXT NOT NULL,
  video_path  VARCHAR(255) DEFAULT NULL,
  created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT fk_music_videos_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE music_videos;