-- drop the table if it exists
DROP TABLE IF EXISTS depositauth;

-- and create the new one
CREATE TABLE depositauth(
   id           INT NOT NULL PRIMARY KEY AUTO_INCREMENT,

   -- user attributes
   employee_id  VARCHAR( 32 ) NOT NULL DEFAULT '',
   computing_id VARCHAR( 32 ) NOT NULL DEFAULT '',
   first_name   VARCHAR( 255 ) NOT NULL DEFAULT '',
   middle_name  VARCHAR( 255 ) NOT NULL DEFAULT '',
   last_name    VARCHAR( 255 ) NOT NULL DEFAULT '',

   -- authorization attributes
   career      VARCHAR( 100 ) NOT NULL DEFAULT '',
   program     VARCHAR( 100 ) NOT NULL DEFAULT '',
   plan        VARCHAR( 100 ) NOT NULL DEFAULT '',
   degree      VARCHAR( 100 ) NOT NULL DEFAULT '',
   title       VARCHAR( 255 ) NOT NULL DEFAULT '',
   doctype     VARCHAR( 100 ) NOT NULL DEFAULT '',

   -- status information
   libra_id    VARCHAR( 100 ) NOT NULL DEFAULT '',
   status      ENUM( 'pending', 'submitted' ) NOT NULL DEFAULT 'pending',
   approved_at TIMESTAMP NULL,
   accepted_at TIMESTAMP NULL,
   exported_at TIMESTAMP NULL,
   created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at  TIMESTAMP NULL

) CHARACTER SET utf8 COLLATE utf8_bin;
-- ) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;