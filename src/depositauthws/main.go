package main

import (
    "fmt"
    "log"
    "net/http"
    "depositauthws/config"
    "depositauthws/dao"
)

func main( ) {

    // access the database

    connectStr := fmt.Sprintf( "%s:%s@tcp(%s)/%s", config.Configuration.DbUser,
        config.Configuration.DbPassphrase, config.Configuration.DbHost, config.Configuration.DbName )

    err := dao.NewDB( connectStr )
    if err != nil {
        log.Fatal( err )
    }

	// setup router and serve...
    router := NewRouter( )
    log.Fatal( http.ListenAndServe( fmt.Sprintf( ":%s", config.Configuration.ServicePort ), router ) )
}

