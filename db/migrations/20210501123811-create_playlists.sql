
-- +migrate Up
CREATE TABLE playlists (
  id              BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id         BIGINT(20) NOT NULL,
  title           VARCHAR(255) NOT NULL,
  created_at      DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at      DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT fk_playlists_user_id FOREIGN KEY (user_id) REFERENCES users (id),
  INDEX(title)
);

-- +migrate Down
DROP TABLE playlists;
