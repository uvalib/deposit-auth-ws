package main

import (
    "net/http"
    "depositauthws/handlers"
    "github.com/gorilla/mux"
)

type Route struct {
   Name        string
   Method      string
   Pattern     string
   HandlerFunc http.HandlerFunc
}

type Routes [] Route

var routes = Routes{

    Route{
        "HealthCheck",
        "GET",
        "/healthcheck",
        handlers.HealthCheck,
    },

    Route{
        "OptionsGet",
        "GET",
        "/options",
        handlers.OptionsGet,
    },

    Route{
        "RegistrationGet",
        "GET",
        "/{id}",
        handlers.RegistrationGet,
    },

    Route{
        "RegistrationSearch",
        "GET",
        "/",
        handlers.RegistrationSearch,
    },

    Route{
        "RegistrationCreate",
        "POST",
        "/",
        handlers.RegistrationCreate,
    },

    Route{
        "RegistrationCreate",
        "OPTIONS",
        "/",
        handlers.RegistrationCreateOptions,
    },

    /*
    Route{
        "RegistrationUpdate",
        "PUT",
        "/{id}",
        handlers.RegistrationUpdate,
    },

    */
    Route{
        "RegistrationDelete",
        "DELETE",
        "/{id}",
        handlers.RegistrationDelete,
    },

}

func NewRouter( ) *mux.Router {

   router := mux.NewRouter().StrictSlash( true )
   for _, route := range routes {

      var handler http.Handler

      handler = route.HandlerFunc
      handler = Logger( handler, route.Name )

      router.
         Methods( route.Method ).
         Path( route.Pattern ).
         Name( route.Name ).
         Handler( handler )
   }

   return router
}
