drop table if exists deposit_auth;
create table deposit_auth (
      id                MEDIUMINT      NOT NULL AUTO_INCREMENT,        /* the primary key, autoincrement */
      eid               VARCHAR( 100 ) NOT NULL DEFAULT '',            /* employee Id */
      cid               VARCHAR( 100 ) NOT NULL DEFAULT '',            /* computing Id */
      first_name        VARCHAR( 255 ) NOT NULL DEFAULT '',            /* student first name */
      middle_name       VARCHAR( 255 ) NOT NULL DEFAULT '',            /* student middle name */
      last_name         VARCHAR( 255 ) NOT NULL DEFAULT '',            /* student last name */

      career            VARCHAR( 100 ) NOT NULL DEFAULT '',            /* student career from SIS */
      program           VARCHAR( 100 ) NOT NULL DEFAULT '',            /* student program from SIS */
      plan              VARCHAR( 100 ) NOT NULL DEFAULT '',            /* student plan from SIS */
      degree            VARCHAR( 100 ) NOT NULL DEFAULT '',            /* student degree from SIS */

      title             VARCHAR( 255 ) NOT NULL DEFAULT '',            /* title of the authorized work */
      doctype           VARCHAR( 100 ) NOT NULL DEFAULT '',            /* document type (milestone in SIS terms) */
      lid               VARCHAR( 100 ) NOT NULL DEFAULT '',            /* document Id (libra Id in SIS terms) */

      approved_at       DATE DEFAULT NULL,                             /* approved by SIS */
      accepted_at       DATETIME DEFAULT NULL,                         /* when deposit was accepted by Libra */
      exported_at       DATETIME DEFAULT NULL,                         /* when exported to SIS */
      created_at        DATETIME DEFAULT NULL,                         /* created_at per Rails */
      updated_at        DATETIME DEFAULT NULL,                         /* updated_at per Rails */

      PRIMARY KEY ( id )
  );