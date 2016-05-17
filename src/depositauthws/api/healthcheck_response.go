package api

type HealthCheckResponse struct {
	CheckType      HealthCheckResult `json:"mysql"`
}

