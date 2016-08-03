package handlers

import (
    "net/http"
    "depositauthws/authtoken"
    "depositauthws/config"
    "depositauthws/dao"
    "depositauthws/sis"
    "depositauthws/logger"
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
        logger.Log( fmt.Sprintf( "ERROR: %s\n", err.Error( ) ) )
        status := http.StatusInternalServerError
        EncodeImportExportResponse( w, status,
            fmt.Sprintf( "%s (%s)", http.StatusText( status ), err ),
            0 )
        return
    }

    // do the export
    err = sis.Exchanger.Export( exports )
    if err != nil {
        logger.Log( fmt.Sprintf( "ERROR: %s\n", err.Error( ) ) )
        status := http.StatusInternalServerError
        EncodeImportExportResponse( w, status,
            fmt.Sprintf( "%s (%s)", http.StatusText( status ), err ),
            0 )
        return
    }

    // update the status so we do not export them again
    err = dao.Database.UpdateExportedDepositAuthorization( exports )
    if err != nil {
        logger.Log( fmt.Sprintf( "ERROR: %s\n", err.Error( ) ) )
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