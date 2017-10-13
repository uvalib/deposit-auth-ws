package api

//
// ImportExportResponse -- response to the import and export requests
//
type ImportExportResponse struct {
   Status  int    `json:"status"`
   Message string `json:"message"`
   Count   int    `json:"count"`
}

//
// end of file
//
