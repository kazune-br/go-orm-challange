-- +migrate Up
CREATE TABLE users (
                       id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                       uid BIGINT NOT NULL,
                       created_at DATETIME  DEFAULT CURRENT_TIMESTAMP NOT NULL ,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL,
                       UNIQUE unique_users_on_uid (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE users;