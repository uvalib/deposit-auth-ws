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

func encodeStandardResponse(w http.ResponseWriter, status int, message string, details []*api.Authorization) {

   logger.Log(fmt.Sprintf("encodeStandardResponse status: %d (%s)\n", status, message))
   jsonAttributes(w)
   coorsAttributes(w)
   w.WriteHeader(status)
   if err := json.NewEncoder(w).Encode(api.StandardResponse{Status: status, Message: message, Details: details}); err != nil {
      log.Fatal(err)
   }
}

func encodeImportResponse(w http.ResponseWriter, status int, message string, newCount int, updateCount int, duplicateCount int, errorCount int ) {

   logger.Log(fmt.Sprintf("encodeImportResponse status: %d (%s)\n", status, message))
   jsonAttributes(w)
   w.WriteHeader(status)
   if err := json.NewEncoder(w).Encode(api.ImportResponse{Status: status,
   Message: message, NewCount: newCount, UpdatedCount: updateCount,
   DuplicateCount: duplicateCount, ErrorCount: errorCount }); err != nil {
      log.Fatal(err)
   }
}

func encodeExportResponse(w http.ResponseWriter, status int, message string, exportCount int, errorCount int ) {

   logger.Log(fmt.Sprintf("encodeExportResponse status: %d (%s)\n", status, message))
   jsonAttributes(w)
   w.WriteHeader(status)
   if err := json.NewEncoder(w).Encode(api.ExportResponse{Status: status, Message: message, ExportCount: exportCount, ErrorCount: errorCount }); err != nil {
      log.Fatal(err)
   }
}

func encodeHealthCheckResponse(w http.ResponseWriter, status int, dbmsg string, importmsg string, exportmsg string) {

   //healthy := status == http.StatusOK
   dbHealthy, importHealthy, exportHealthy := true, true, true
   if len(dbmsg) != 0 {
      dbHealthy = false
   }
   if len(importmsg) != 0 {
      importHealthy = false
   }
   if len(exportmsg) != 0 {
      exportHealthy = false
   }
   jsonAttributes(w)
   w.WriteHeader(status)
   if err := json.NewEncoder(w).Encode(api.HealthCheckResponse{
      DbCheck:       api.HealthCheckResult{Healthy: dbHealthy, Message: dbmsg},
      ImportFsCheck: api.HealthCheckResult{Healthy: importHealthy, Message: importmsg},
      ExportFsCheck: api.HealthCheckResult{Healthy: exportHealthy, Message: exportmsg}}); err != nil {
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
   if err := json.NewEncoder(w).Encode(api.RuntimeResponse{Version: version, CPUCount: cpus, GoRoutineCount: goroutines, ObjectCount: heapcount, AllocatedMemory: alloc}); err != nil {
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

func isEmpty(param string) bool {
   return len(strings.TrimSpace(param)) == 0
}

// map any field values as necessary
func mapResultsFieldValues(details []*api.Authorization) {
   for _, d := range details {
      d.Department, _ = mapper.MapField("department", d.Plan)
      d.Degree, _ = mapper.MapField("degree", d.Degree)
   }
}

//
// end of file
//
