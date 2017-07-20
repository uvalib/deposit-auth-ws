package main

import (
	"depositauthws/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"HealthCheck",
		"GET",
		"/healthcheck",
		handlers.HealthCheck,
	},

	Route{
		"VersionInfo",
		"GET",
		"/version",
		handlers.VersionInfo,
	},

	Route{
		"RuntimeInfo",
		"GET",
		"/runtime",
		handlers.RuntimeInfo,
	},

	Route{
		"AuthorizationGet",
		"GET",
		"/{id}",
		handlers.AuthorizationGet,
	},

	Route{
		"AuthorizationSearch",
		"GET",
		"/",
		handlers.AuthorizationSearch,
	},

	Route{
		"AuthorizationFulfill",
		"PUT",
		"/{id}",
		handlers.AuthorizationFulfill,
	},

	Route{
		"AuthorizationImport",
		"POST",
		"/import",
		handlers.AuthorizationImport,
	},

	Route{
		"AuthorizationExport",
		"POST",
		"/export",
		handlers.AuthorizationExport,
	},

	Route{
		"AuthorizationDelete",
		"DELETE",
		"/{id}",
		handlers.AuthorizationDelete,
	},
}

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
