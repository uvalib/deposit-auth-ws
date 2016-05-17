package config

import (
    "flag"
    "log"
)

type Config struct {
    ServicePort        string
    DbHost             string
    DbName             string
    DbUser             string
    DbPassphrase       string
    AuthTokenEndpoint  string
    Debug              bool
}

var Configuration = LoadConfig( )

func LoadConfig( ) Config {

    c := Config{}

    // process command line flags and setup configuration
    flag.StringVar( &c.ServicePort, "port", "8080", "The service listen port" )
    flag.StringVar( &c.DbHost, "dbhost", "mysqldev.lib.virginia.edu:3306", "The database server hostname:port" )
    flag.StringVar( &c.DbName, "dbname", "depositreg_development", "The database name" )
    flag.StringVar( &c.DbUser, "dbuser", "depositreg", "The database username" )
    flag.StringVar( &c.DbPassphrase, "dbpassword", "dbpassword", "The database passphrase")
    flag.StringVar( &c.AuthTokenEndpoint, "tokenauth", "http://docker1.lib.virginia.edu:8200", "The token authentication endpoint")
    flag.BoolVar( &c.Debug, "debug", false, "Enable debugging")

    flag.Parse()

    log.Printf( "ServicePort:       %s", c.ServicePort )
    log.Printf( "DbHost:            %s", c.DbHost )
    log.Printf( "DbName:            %s", c.DbName )
    log.Printf( "DbUser:            %s", c.DbUser )
    log.Printf( "DbPassphrase:      %s", c.DbPassphrase )
    log.Printf( "AuthTokenEndpoint  %s", c.AuthTokenEndpoint )
    log.Printf( "Debug              %t", c.Debug )

    return c
}

