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
      "HealthCheck",
      "GET",
      "/healthcheck",
      handlers.HealthCheck,
   },

   route{
      "VersionInfo",
      "GET",
      "/version",
      handlers.VersionInfo,
   },

   route{
      "RuntimeInfo",
      "GET",
      "/runtime",
      handlers.RuntimeInfo,
   },

   route{
      "AuthorizationGet",
      "GET",
      "/{id}",
      handlers.AuthorizationGet,
   },

   route{
      "AuthorizationSearch",
      "GET",
      "/",
      handlers.AuthorizationSearch,
   },

   route{
      "AuthorizationFulfill",
      "PUT",
      "/{id}",
      handlers.AuthorizationFulfill,
   },

   route{
      "AuthorizationImport",
      "POST",
      "/import",
      handlers.AuthorizationImport,
   },

   route{
      "AuthorizationExport",
      "POST",
      "/export",
      handlers.AuthorizationExport,
   },

   route{
      "AuthorizationDelete",
      "DELETE",
      "/{id}",
      handlers.AuthorizationDelete,
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