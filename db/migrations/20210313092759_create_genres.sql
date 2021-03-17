-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE genres (
  id              BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name            VARCHAR(255) NOT NULL,
  created_at      DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at      DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX(name)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE genres;
