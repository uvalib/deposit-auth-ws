package config

import (
	"flag"
	"fmt"
	"github.com/uvalib/deposit-auth-ws/depositauthws/logger"
	"strings"
)

//
// Config -- our configuration structure
//
type Config struct {
	ServicePort       string // our listen port
	DbSecure          string // do we use TLS
	DbHost            string // hostname of database server
	DbName            string // database name
	DbUser            string // database user name
	DbPassphrase      string // database user password
	DbTimeout         string // connection/read/write timeout
	ImportFs          string
	ExportFs          string
	SharedSecret      string
	Debug             bool
}

//
// Configuration -- our configuration instance
//
var Configuration = loadConfig()

func loadConfig() Config {

	// default value for the database timeout
	c := Config{DbTimeout: "15s"}

	// process command line flags and setup configuration
	flag.StringVar(&c.ServicePort, "port", "8080", "The service listen port")
	flag.StringVar(&c.DbSecure, "dbsecure", "false", "Use TLS for the database connection")
	flag.StringVar(&c.DbHost, "dbhost", "mysqldev.lib.virginia.edu:3306", "The database server hostname:port")
	flag.StringVar(&c.DbName, "dbname", "depositauth_development", "The database name")
	flag.StringVar(&c.DbUser, "dbuser", "depositauth", "The database username")
	flag.StringVar(&c.DbPassphrase, "dbpassword", "", "The database passphrase")
	flag.StringVar(&c.ImportFs, "importfs", "/tmp/import", "The import filesystem")
	flag.StringVar(&c.ExportFs, "exportfs", "/tmp/export", "The export filesystem")
	flag.StringVar(&c.SharedSecret, "secret", "", "The JWT shared secret")
	flag.BoolVar(&c.Debug, "debug", false, "Enable debugging")

	flag.Parse()

	logger.Log(fmt.Sprintf("ServicePort:       %s", c.ServicePort))
	logger.Log(fmt.Sprintf("DbSecure:          %s", c.DbSecure))
	logger.Log(fmt.Sprintf("DbHost:            %s", c.DbHost))
	logger.Log(fmt.Sprintf("DbName:            %s", c.DbName))
	logger.Log(fmt.Sprintf("DbUser:            %s", c.DbUser))
	logger.Log(fmt.Sprintf("DbPassphrase:      %s", strings.Repeat("*", len(c.DbPassphrase))))
	logger.Log(fmt.Sprintf("DbTimeout:         %s", c.DbTimeout))
	logger.Log(fmt.Sprintf("ImportFs:          %s", c.ImportFs))
	logger.Log(fmt.Sprintf("ExportFs:          %s", c.ExportFs))
	logger.Log(fmt.Sprintf("SharedSecret:      %s", strings.Repeat("*", len(c.SharedSecret))))
	logger.Log(fmt.Sprintf("Debug:             %t", c.Debug))

	return c
}

//
// end of file
//
