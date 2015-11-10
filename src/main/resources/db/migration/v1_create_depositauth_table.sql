drop table if exists deposit_auth;
create table deposit_auth (
      id      MEDIUMINT      NOT NULL AUTO_INCREMENT,
      cid     VARCHAR( 100 ) NOT NULL DEFAULT '',
      doctype VARCHAR( 100 ) NOT NULL DEFAULT '',
      lid     VARCHAR( 100 ) NOT NULL DEFAULT '',
      PRIMARY KEY ( id )
  );