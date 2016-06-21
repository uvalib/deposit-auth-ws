package handlers

import (
    "log"
    "fmt"
    "net/http"
    "depositauthws/authtoken"
    "depositauthws/config"
    "depositauthws/sis"
    "depositauthws/dao"
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

    // import each record and keep track of progress
    okCount := 0
    duplicateCount := 0
    errorCount := 0
    for _, e := range imports {

        // check to see if this record already exists
        exists, err := dao.Database.DepositAuthorizationExists( *e )
        if err != nil {
            errorCount += 1
        } else {
           if exists == true {
            duplicateCount += 1
            log.Printf( "record already exists, ignoring (%s/%s/%s/%s)", e.ComputingId, e.Degree, e.Plan, e.Title )
           } else {
               _, err = dao.Database.CreateDepositAuthorization( *e )
               if err != nil {
                   log.Printf( "Error inserting record; ignoring %s for (%s/%s/%s/%s)", err, e.ComputingId, e.Degree, e.Plan, e.Title )
                   errorCount += 1
               } else {
                   okCount += 1
               }
           }
        }
    }

    // did we encounter any errors
    if errorCount != 0 {
        status := http.StatusCreated
        EncodeImportExportResponse( w, status,
            fmt.Sprintf( "%s (%d errors encountered)", http.StatusText( status ), errorCount ),
            okCount )
        return
    }

    // its all over
    status := http.StatusOK
    EncodeImportExportResponse( w, status, http.StatusText( status ), okCount )
}