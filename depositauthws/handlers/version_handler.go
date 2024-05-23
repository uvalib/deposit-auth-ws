package handlers

import (
	"net/http"
)

// VersionHandler - get version handler
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	encodeVersionResponse(w, http.StatusOK, Version())
}

//
// end of file
//
