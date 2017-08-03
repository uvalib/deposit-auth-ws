package handlers

import (
	"depositauthws/api"
	"depositauthws/logger"
	"depositauthws/mapper"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func EncodeStandardResponse(w http.ResponseWriter, status int, message string, details []*api.Authorization) {

	logger.Log(fmt.Sprintf("EncodeStandardResponse status: %d (%s)\n", status, message))
	jsonAttributes(w)
	coorsAttributes(w)
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(api.StandardResponse{Status: status, Message: message, Details: details}); err != nil {
		log.Fatal(err)
	}
}

func EncodeImportExportResponse(w http.ResponseWriter, status int, message string, count int) {

	logger.Log(fmt.Sprintf("EncodeImportExportResponse status: %d (%s)\n", status, message))
	jsonAttributes(w)
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(api.ImportExportResponse{Status: status, Message: message, Count: count}); err != nil {
		log.Fatal(err)
	}
}

func EncodeHealthCheckResponse(w http.ResponseWriter, status int, dbmsg string, importmsg string, exportmsg string) {

	//healthy := status == http.StatusOK
	db_healthy, import_healthy, export_healthy := true, true, true
	if len(dbmsg) != 0 {
		db_healthy = false
	}
	if len(importmsg) != 0 {
		import_healthy = false
	}
	if len(exportmsg) != 0 {
		export_healthy = false
	}
	jsonAttributes(w)
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(api.HealthCheckResponse{
		DbCheck:       api.HealthCheckResult{Healthy: db_healthy, Message: dbmsg},
		ImportFsCheck: api.HealthCheckResult{Healthy: import_healthy, Message: importmsg},
		ExportFsCheck: api.HealthCheckResult{Healthy: export_healthy, Message: exportmsg}}); err != nil {
		log.Fatal(err)
	}
}

func encodeVersionResponse(w http.ResponseWriter, status int, version string) {
	jsonAttributes(w)
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(api.VersionResponse{Version: version}); err != nil {
		log.Fatal(err)
	}
}

func encodeRuntimeResponse(w http.ResponseWriter, status int, version string, cpus int, goroutines int, heapcount uint64, alloc uint64) {
	jsonAttributes(w)
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(api.RuntimeResponse{Version: version, CpuCount: cpus, GoRoutineCount: goroutines, ObjectCount: heapcount, AllocatedMemory: alloc}); err != nil {
		log.Fatal(err)
	}
}

func jsonAttributes(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func coorsAttributes(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func NotEmpty(param string) bool {
	return len(strings.TrimSpace(param)) != 0
}

// map any field values as necessary
func MapResultsFieldValues(details []*api.Authorization) {
	for _, d := range details {
		d.Department, _ = mapper.MapField("department", d.Plan)
		d.Degree, _ = mapper.MapField("degree", d.Degree)
	}
}
