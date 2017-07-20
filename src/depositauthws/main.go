package main

import (
	"depositauthws/config"
	"depositauthws/dao"
	"depositauthws/handlers"
	"depositauthws/logger"
	"depositauthws/mapper"
	"depositauthws/sis"
	"fmt"
	"log"
	"net/http"
)

func main() {

	logger.Log(fmt.Sprintf("===> version: '%s' <===", handlers.Version()))

	// access the database
	connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowOldPasswords=1&strict=true&sql_notes=false", config.Configuration.DbUser,
		config.Configuration.DbPassphrase, config.Configuration.DbHost, config.Configuration.DbName)

	err := dao.NewDB(connectStr)
	if err != nil {
		log.Fatal(err)
	}

	// the filesystem used for SIS exchange
	err = sis.NewExchanger(config.Configuration.ImportFs, config.Configuration.ExportFs)
	if err != nil {
		log.Fatal(err)
	}

	// the mapping cache for various fields
	err = mapper.LoadMappingCache()
	if err != nil {
		log.Fatal(err)
	}

	// setup router and serve...
	router := NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Configuration.ServicePort), router))
}
