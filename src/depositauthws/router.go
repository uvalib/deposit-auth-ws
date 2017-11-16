package main

import (
	"depositauthws/handlers"
	"github.com/gorilla/mux"
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
		"RuntimeHandler",
		"GET",
		"/runtime",
		handlers.RuntimeHandler,
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
	for _, route := range routes {

		var handler http.Handler

		handler = route.HandlerFunc
		handler = HandlerLogger(handler, route.Name)

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
