-- drop the table if it exists
DROP TABLE IF EXISTS inbound;

-- and create the new one
CREATE TABLE inbound(
   id           INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
   deposit_id   INT NOT NULL,
   created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

) CHARACTER SET utf8 COLLATE utf8_bin;
-- ) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;