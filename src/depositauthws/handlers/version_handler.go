package handlers

import (
    "net/http"
)

func VersionGet( w http.ResponseWriter, r *http.Request ) {
    encodeVersionResponse( w, http.StatusOK, Version( ) )
}