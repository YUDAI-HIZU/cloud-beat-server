-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE user_device_tokens (
  user_id      BIGINT(20) NOT NULL PRIMARY KEY,
  device_token TEXT NOT NULL,
  created_at   DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at   DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user_device_tokens;