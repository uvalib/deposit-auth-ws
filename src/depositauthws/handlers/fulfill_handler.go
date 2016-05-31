package handlers

import (
    "net/http"
    "github.com/gorilla/mux"
    "depositauthws/authtoken"
    "depositauthws/config"
    "depositauthws/dao"
    "log"
    "fmt"
)

func AuthorizationFulfill( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    id := vars[ "id" ]
    token := r.URL.Query( ).Get( "auth" )
    did := r.URL.Query( ).Get( "deposit" )

    // parameters OK ?
    if NotEmpty( id ) == false || NotEmpty( token ) == false || NotEmpty( did ) == false {
        status := http.StatusBadRequest
        EncodeStandardResponse( w, status, http.StatusText( status ), nil )
        return
    }

    // validate the token
    if authtoken.Validate( config.Configuration.AuthTokenEndpoint, token ) == false {
        status := http.StatusForbidden
        EncodeStandardResponse( w, status, http.StatusText( status ), nil )
        return
    }

    // get the authorization details
    reqs, err := dao.Database.GetDepositAuthorizationById( id )
    if err != nil {
        log.Println( err )
        status := http.StatusInternalServerError
        EncodeStandardResponse( w, status,
            fmt.Sprintf( "%s (%s)", http.StatusText( status ), err ),
            nil )
        return
    }

    // we did not find the item, return 404
    if reqs == nil || len( reqs ) == 0 {
        status := http.StatusNotFound
        EncodeStandardResponse( w, status, http.StatusText( status ), nil )
        return
    }

    // handle the fulfill
    err = dao.Database.UpdateFulfilledDepositAuthorization( id, did )
    if err != nil {
        log.Println( err )
        status := http.StatusInternalServerError
        EncodeStandardResponse( w, status,
            fmt.Sprintf( "%s (%s)", http.StatusText( status ), err ),
            nil )
        return
    }

    // its all over
    status := http.StatusOK
    EncodeStandardResponse( w, status, http.StatusText( status ), nil )
}