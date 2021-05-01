
-- +migrate Up
CREATE TABLE external_links (
  user_id     BIGINT(20) NOT NULL PRIMARY KEY,
  twitter     VARCHAR(255) DEFAULT NULL,
  sound_cloud VARCHAR(255) DEFAULT NULL,
  facebook    VARCHAR(255) DEFAULT NULL,
  youtube     VARCHAR(255) DEFAULT NULL,
  instagram   VARCHAR(255) DEFAULT NULL,
  created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT fk_external_links_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE external_links;
