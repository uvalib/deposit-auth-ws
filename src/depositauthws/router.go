package main

import (
	"depositauthws/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routeSlice []route

var routes = routeSlice{

	route{
		"HealthCheckHandler",
		"GET",
		"/healthcheck",
		handlers.HealthCheckHandler,
	},

	route{
		"VersionHandler",
		"GET",
		"/version",
		handlers.VersionHandler,
	},

	route{
		"InboundHandler",
		"GET",
		"/inbound",
		handlers.InboundHandler,
	},

	route{
		"GetHandler",
		"GET",
		"/{id}",
		handlers.GetHandler,
	},

	route{
		"SearchHandler",
		"GET",
		"/",
		handlers.SearchHandler,
	},

	route{
		"FulfillHandler",
		"PUT",
		"/{id}",
		handlers.FulfillHandler,
	},

	route{
		"ImportHandler",
		"POST",
		"/import",
		handlers.ImportHandler,
	},

	route{
		"ExportHandler",
		"POST",
		"/export",
		handlers.ExportHandler,
	},

	route{
		"DeleteHandler",
		"DELETE",
		"/{id}",
		handlers.DeleteHandler,
	},
}

//
// NewRouter -- build and return the router
//
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	// add the route for the prometheus metrics
	router.Handle("/metrics", HandlerLogger(promhttp.Handler(), "promhttp.Handler"))

	for _, route := range routes {

		var handler http.Handler = route.HandlerFunc
		handler = HandlerLogger(handler, route.Name)
		handler = prometheus.InstrumentHandler(route.Name, handler)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

//
// end of file
//
