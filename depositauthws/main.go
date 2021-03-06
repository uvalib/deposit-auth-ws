package main

import (
	"fmt"
	"github.com/uvalib/deposit-auth-ws/depositauthws/config"
	"github.com/uvalib/deposit-auth-ws/depositauthws/dao"
	"github.com/uvalib/deposit-auth-ws/depositauthws/handlers"
	"github.com/uvalib/deposit-auth-ws/depositauthws/logger"
	"github.com/uvalib/deposit-auth-ws/depositauthws/mapper"
	"github.com/uvalib/deposit-auth-ws/depositauthws/sis"
	"log"
	"net/http"
	"time"
)

func main() {

	logger.Log(fmt.Sprintf("===> version: '%s' <===", handlers.Version()))

	// create the storage singleton
	err := dao.NewDatastore()
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
