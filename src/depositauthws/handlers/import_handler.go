package handlers

import (
    "log"
    "fmt"
    "net/http"
    "depositauthws/authtoken"
    "depositauthws/config"
    "depositauthws/sis"
//    "depositauthws/dao"
)

func AuthorizationImport( w http.ResponseWriter, r *http.Request ) {

    token := r.URL.Query( ).Get( "auth" )

    // parameters OK ?
    if NotEmpty( token ) == false {
        status := http.StatusBadRequest
        EncodeImportExportResponse( w, status, http.StatusText( status ), 0 )
        return
    }

    // validate the token
    if authtoken.Validate( config.Configuration.AuthTokenEndpoint, token ) == false {
        status := http.StatusForbidden
        EncodeImportExportResponse( w, status, http.StatusText( status ), 0 )
        return
    }

    // get the details ready to be imported
    imports, err := sis.Exchanger.Import( )
    if err != nil {
        log.Println( err )
        status := http.StatusInternalServerError
        EncodeImportExportResponse( w, status,
            fmt.Sprintf( "%s (%s)", http.StatusText( status ), err ),
            0 )
        return
    }

    // if we have nothing to import, bail out
    if imports == nil || len( imports ) == 0 {
        status := http.StatusOK
        EncodeImportExportResponse( w, status, http.StatusText( status ), 0 )
        return
    }

    // its all over
    status := http.StatusOK
    EncodeImportExportResponse( w, status, http.StatusText( status ), len( imports ) )
}