drop table if exists deposit_auth;
create table deposit_auth (
      id                MEDIUMINT      NOT NULL AUTO_INCREMENT,
      cid               VARCHAR( 100 ) NOT NULL DEFAULT '',
      doctype           VARCHAR( 100 ) NOT NULL DEFAULT '',
      lid               VARCHAR( 100 ) NOT NULL DEFAULT '',
      title             VARCHAR( 255 ) NOT NULL DEFAULT '',
      program           VARCHAR( 100 ) NOT NULL DEFAULT '',

      approved_at       DATE DEFAULT NULL,
      exported_at       DATETIME DEFAULT NULL,
      created_at        DATETIME DEFAULT NULL,
      updated_at        DATETIME DEFAULT NULL,

      PRIMARY KEY ( id )
  );