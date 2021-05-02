
-- +migrate Up
CREATE TABLE users (
  id                 BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  uid                VARCHAR(255) UNIQUE NOT NULL,
  icon_name          VARCHAR(255) DEFAULT NULL,
  display_name       VARCHAR(255) DEFAULT NULL,
  web_url            VARCHAR(255) DEFAULT NULL,
  introduction       VARCHAR(255) DEFAULT NULL,
  created_at         DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at         DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX(uid, display_name, introduction)
);

-- +migrate Down
DROP TABLE users;