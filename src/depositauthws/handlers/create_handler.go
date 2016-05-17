package handlers

import (
    "log"
    "fmt"
    "strings"
    "net/http"
    "encoding/json"
//    "github.com/gorilla/mux"
    "depositauthws/authtoken"
    "depositauthws/config"
    "depositauthws/dao"
    "depositauthws/api"
)

func RegistrationCreate( w http.ResponseWriter, r *http.Request ) {

    token := r.URL.Query( ).Get( "auth" )

    // parameters OK ?
    if NotEmpty( token ) == false {
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

    decoder := json.NewDecoder( r.Body )
    reg := api.Registration{ }

    if err := decoder.Decode( &reg ); err != nil {
        status := http.StatusBadRequest
        EncodeStandardResponse( w, status, http.StatusText( status ), nil )
        return
    }

    // create results list
    results := make([ ] * api.Registration, 0 )

    // split the user list of appropriate
    users := strings.Split( reg.For, "," )

    for _, u := range users {

        reg.For = strings.TrimSpace( u )
        rg, err := dao.Database.CreateDepositRequest( reg )
        if err != nil {
            log.Println(err)
            status := http.StatusInternalServerError
            EncodeStandardResponse(w, status,
                fmt.Sprintf("%s (%s)", http.StatusText( status ), err),
                nil)
            return
        }

        results = append(results, rg)
    }

    status := http.StatusOK
    EncodeStandardResponse( w, status, http.StatusText( status ), results )
}

func RegistrationCreateOptions( w http.ResponseWriter, r *http.Request ) {
    w.Header( ).Add( "Access-Control-Allow-Methods", "POST" )
    EncodeStandardResponse( w, http.StatusOK, http.StatusText( http.StatusOK ), nil )
}