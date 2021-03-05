-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE images (
  id           BIGINT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name         VARCHAR(255) NOT NULL,
  owner_id     BIGINT(20) NOT NULL,
  owner_type   VARCHAR(255) NOT NULL,
  created_at   DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at   DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE images;