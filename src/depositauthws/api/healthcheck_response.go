package api

type HealthCheckResponse struct {
	DbCheck        HealthCheckResult `json:"mysql"`
    ImportFsCheck  HealthCheckResult `json:"import_fs"`
    ExportFsCheck  HealthCheckResult `json:"export_fs"`
}

