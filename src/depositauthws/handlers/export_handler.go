package handlers

import (
    "net/http"
    "depositauthws/authtoken"
    "depositauthws/config"
    "depositauthws/dao"
    "depositauthws/sis"
    "log"
    "fmt"
)

func AuthorizationExport( w http.ResponseWriter, r *http.Request ) {

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

    // get the details ready to be exported
    exports, err := dao.Database.GetDepositAuthorizationForExport( )
    if err != nil {
        log.Println( err )
        status := http.StatusInternalServerError
        EncodeImportExportResponse( w, status,
            fmt.Sprintf( "%s (%s)", http.StatusText( status ), err ),
            0 )
        return
    }

    // if we have nothing to export, bail out
    if exports == nil || len( exports ) == 0 {
        status := http.StatusOK
        EncodeImportExportResponse( w, status, http.StatusText( status ), 0 )
        return
    }

    // do the export
    err = sis.Exchanger.Export( exports )
    if err != nil {
        log.Println( err )
        status := http.StatusInternalServerError
        EncodeImportExportResponse( w, status,
            fmt.Sprintf( "%s (%s)", http.StatusText( status ), err ),
            0 )
        return
    }

    // update the status so we do not export them again
    err = dao.Database.UpdatedExportedDepositAuthorization( exports )
    if err != nil {
        log.Println( err )
        status := http.StatusInternalServerError
        EncodeImportExportResponse( w, status,
            fmt.Sprintf( "%s (%s)", http.StatusText( status ), err ),
            0 )
        return
    }

    // its all over
    status := http.StatusOK
    EncodeImportExportResponse( w, status, http.StatusText( status ), len( exports ) )
}