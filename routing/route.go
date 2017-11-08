// SlackWiener backend routing package
// Author: Teemu Poikkijoki <poidokoff@gmail.com>

// package routing provides routing utilities for slackwiener backend
package routing

import (
  "net/http"
  "strings"
  "github.com/gorilla/mux"
  "bitbucket.org/ncolabs/slackwiener_backend/logging"
)

// EndpointHandlerFunc contains a specification for an Endpoint handler that provides a normal http handler callback 
// and optional authorization
type EndpointHandlerFunc struct {
  Function func(map[string]string, http.ResponseWriter, *http.Request)
}


// HandlerFunctions maps named endpoint functions to a handler func
// Example: "ninjaIndex", {function: ninjas.GetNinjas}
type HandlerFunctions map[string]EndpointHandlerFunc


// HandlerRoutes maps named routes to handler function specs
// Example: "GET /ninjas/", "ninjaIndex"
type HandlerRoutes map[string]string


// CreaterRouter parses routes and maps them to provided handler functions, returning a new mux.Router
func CreateRouter(routes HandlerRoutes, funcs HandlerFunctions) *mux.Router {

  var routeList []string
  var r = mux.NewRouter()

  // Iterate all routes and split them into VERB route
  for k, v := range routes {
    split := strings.SplitAfterN(k, " ", 3)

    split[0] = strings.Trim(split[0], " ")
    split[1] = strings.Trim(split[1], " ")

    if !validateVerb(split[0]) {
      logging.Log.Errorf("Unknown HTTP verb: %s", split[0])
      panic("UNKNOWN HTTP VERB!")
    }

    routeList = append(routeList, split[0] + "\t" + split[1] + "\t=>" + v)

    func(name string) {
      handler := funcs[name]
      if handler.Function == nil {
        logging.Log.Errorf("Function not found for action %s", name)
        panic("Route handler not found")
      }
      logging.Log.Info("r.HandleFunc(%s, _toDispatcherFunc(handler.Function)).Methods(%s)", split[1], split[0])
      r.HandleFunc(split[1], _toDispatcherFunc(handler.Function)).Methods(split[0])
    }(v)
  }

  logging.Log.Info("Registered routes: " + strings.Join(routeList, "\n\t"))
  return r
}


// _toDispatcherFunc creates a decorator to pass request params as a mapping to our callbacks
func _toDispatcherFunc(dispatcher func(map[string]string, http.ResponseWriter, *http.Request)) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    dispatcher(params, w, req)
  })
}


// validateVerb validates that the provided verb is an accepted HTTP action
func validateVerb(verb string) bool {
  return verb == "GET" || verb == "PUT" || verb == "POST" || verb == "DELETE" || verb == "HEAD" || verb == "OPTIONS"
}

