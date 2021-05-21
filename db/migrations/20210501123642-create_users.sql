
-- +migrate Up
CREATE TABLE users (
  id                 VARCHAR(255) UNIQUE NOT NULL PRIMARY KEY,
  icon_name          VARCHAR(255) DEFAULT NULL,
  display_name       VARCHAR(255) DEFAULT NULL,
  web_url            VARCHAR(255) DEFAULT NULL,
  introduction       VARCHAR(255) DEFAULT NULL,
  created_at         DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at         DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE users;
