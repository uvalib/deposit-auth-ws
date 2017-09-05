package handlers

import (
	"net/http"
	"runtime"
)

func RuntimeInfo(w http.ResponseWriter, r *http.Request) {

	version := runtime.Version()
	ncpu := runtime.NumCPU()
	ngr := runtime.NumGoroutine()
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)
	encodeRuntimeResponse(w, http.StatusOK, version, ncpu, ngr, m.HeapObjects, m.Alloc)
}