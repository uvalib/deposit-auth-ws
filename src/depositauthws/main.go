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
	"time"
)

func main() {

	logger.Log(fmt.Sprintf("===> version: '%s' <===", handlers.Version()))

	// access the database
	timeout := "10s" // connect, read and write timeout in seconds
	connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowOldPasswords=1&sql_notes=false&timeout=%s&readTimeout=%s&writeTimeout=%s",
		config.Configuration.DbUser,
		config.Configuration.DbPassphrase,
		config.Configuration.DbHost,
		config.Configuration.DbName,
		timeout, timeout, timeout)

	err := dao.NewDB(connectStr)
	if err != nil {
		log.Fatal(err)
	}

	// the filesystem used for sisImpliementation exchange
	err = sis.NewExchanger(config.Configuration.ImportFs, config.Configuration.ExportFs)
	if err != nil {
		log.Fatal(err)
	}

	// the mapping cache for various fields
	err = mapper.LoadMappingCache()
	if err != nil {
		log.Fatal(err)
	}

	// setup router and server...
	serviceTimeout := 15 * time.Second
	router := NewRouter()
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Configuration.ServicePort),
		Handler:      router,
		ReadTimeout:  serviceTimeout,
		WriteTimeout: serviceTimeout,
	}
	log.Fatal(server.ListenAndServe())
}

//
// end of file
//
