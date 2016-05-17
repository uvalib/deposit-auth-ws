package handlers

import (
    "log"
    "fmt"
    "net/http"
    //"depositauthws/authtoken"
    //"depositauthws/config"
    "depositauthws/api"
    "depositauthws/dao"
)

func OptionsGet( w http.ResponseWriter, r *http.Request ) {

    //token := r.URL.Query( ).Get( "auth" )

    // parameters OK ?
    //if NotEmpty( token ) == false {
    //    EncodeOptionsResponse( w, http.StatusBadRequest, http.StatusText( http.StatusBadRequest ), nil )
    //    return
    //}

    // validate the token
    //if authtoken.Validate( config.Configuration.AuthTokenEndpoint, token ) == false {
    //    EncodeOptionsResponse( w, http.StatusForbidden, http.StatusText( http.StatusForbidden ), nil )
    //    return
    //}

    departments, err := dao.Database.GetFieldSet( "department" )
    if err != nil {
        log.Println( err )
        status := http.StatusInternalServerError
        EncodeOptionsResponse( w, status,
            fmt.Sprintf( "%s (%s)", http.StatusText( status ), err ),
            nil )
        return
    }

    degrees, err := dao.Database.GetFieldSet( "degree" )
    if err != nil {
        log.Println( err )
        status := http.StatusInternalServerError
        EncodeOptionsResponse( w, status,
            fmt.Sprintf( "%s (%s)", http.StatusText( status ), err ),
            nil )
        return
    }

    options := api.Options{ Department: departments, Degree: degrees }

    status := http.StatusOK
    EncodeOptionsResponse( w, status, http.StatusText( status ), &options )
}