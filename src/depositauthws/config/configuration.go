package config

import (
	"depositauthws/logger"
	"flag"
	"fmt"
	"strings"
)

//
// Config -- our configuration structure
//
type Config struct {
	ServicePort       string
	DbHost            string
	DbName            string
	DbUser            string
	DbPassphrase      string
	AuthTokenEndpoint string
	ImportFs          string
	ExportFs          string
	Timeout           int
	Debug             bool
}

//
// Configuration -- our configuration instance
//
var Configuration = loadConfig()

func loadConfig() Config {

	c := Config{}

	// process command line flags and setup configuration
	flag.StringVar(&c.ServicePort, "port", "8080", "The service listen port")
	flag.StringVar(&c.DbHost, "dbhost", "mysqldev.lib.virginia.edu:3306", "The database server hostname:port")
	flag.StringVar(&c.DbName, "dbname", "depositauth_development", "The database name")
	flag.StringVar(&c.DbUser, "dbuser", "depositauth", "The database username")
	flag.StringVar(&c.DbPassphrase, "dbpassword", "", "The database passphrase")
	flag.StringVar(&c.ImportFs, "importfs", "/tmp/import", "The import filesystem")
	flag.StringVar(&c.ExportFs, "exportfs", "/tmp/export", "The export filesystem")
	flag.StringVar(&c.AuthTokenEndpoint, "tokenauth", "http://docker1.lib.virginia.edu:8200", "The token authentication endpoint")
	flag.IntVar(&c.Timeout, "timeout", 15, "The external service timeout in seconds")
	flag.BoolVar(&c.Debug, "debug", false, "Enable debugging")

	flag.Parse()

	logger.Log(fmt.Sprintf("ServicePort:       %s", c.ServicePort))
	logger.Log(fmt.Sprintf("DbHost:            %s", c.DbHost))
	logger.Log(fmt.Sprintf("DbName:            %s", c.DbName))
	logger.Log(fmt.Sprintf("DbUser:            %s", c.DbUser))
	logger.Log(fmt.Sprintf("DbPassphrase:      %s", strings.Repeat("*", len(c.DbPassphrase))))
	logger.Log(fmt.Sprintf("AuthTokenEndpoint  %s", c.AuthTokenEndpoint))
	logger.Log(fmt.Sprintf("ImportFs           %s", c.ImportFs))
	logger.Log(fmt.Sprintf("ExportFs           %s", c.ExportFs))
	logger.Log(fmt.Sprintf("Timeout:           %d", c.Timeout))
	logger.Log(fmt.Sprintf("Debug              %t", c.Debug))

	return c
}

//
// end of file
//
