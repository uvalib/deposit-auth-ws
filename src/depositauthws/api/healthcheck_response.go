package api

//
// HealthCheckResponse -- response to the health check query
//
type HealthCheckResponse struct {
   DbCheck       HealthCheckResult `json:"mysql"`
   ImportFsCheck HealthCheckResult `json:"import_fs"`
   ExportFsCheck HealthCheckResult `json:"export_fs"`
}

//
// end of file
//
