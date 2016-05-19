package api

type ImportExportResponse struct {
   Status        int           `json:"status"`
   Message       string        `json:"message"`
   Count         int           `json:"count"`
}

