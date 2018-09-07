package api

//
// ExportResponse -- response to the export request
//
type ExportResponse struct {
	Status      int    `json:"status"`
	Message     string `json:"message"`
	ExportCount int    `json:"export_count"`
	ErrorCount  int    `json:"error_count"`
}

//
// end of file
//
