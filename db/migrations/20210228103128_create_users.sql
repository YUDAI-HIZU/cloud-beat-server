-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
  id                 BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  uid                VARCHAR(255) UNIQUE NOT NULL,
  icon_path          VARCHAR(255) DEFAULT NULL,
  cover_path         VARCHAR(255) DEFAULT NULL,
  display_name       VARCHAR(255) DEFAULT NULL,
  web_url            VARCHAR(255) DEFAULT NULL,
  twitter            VARCHAR(255) DEFAULT NULL,
  sound_cloud        VARCHAR(255) DEFAULT NULL,
  facebook           VARCHAR(255) DEFAULT NULL,
  youtube            VARCHAR(255) DEFAULT NULL,
  instagram          VARCHAR(255) DEFAULT NULL,
  introduction       VARCHAR(255) DEFAULT NULL,
  created_at         DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at         DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX(uid, display_name, introduction)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
