// Package events provides an API similar to routing in order to connect different kinds of events their proper handlers
package events

import (
  "net/http"
  "fmt"
  "bitbucket.org/ncolabs/slackwiener_backend/logging"
)

// SlackEventHandler describes a handler for a slack event. Function is the callback
type SlackEventHandlerFunc struct {
  Function      func(map[string]string, http.ResponseWriter, *http.Request)
}


// SlackEventHandler maps an event by type to a corresponding handler. The map index is a slack event type
type SlackHandlerFunctions map[string]SlackEventHandlerFunc


type SlackEventHandler struct {
  Handlers  SlackHandlerFunctions
}

var dp *SlackEventHandler

func InitializeDispatcher(handlers SlackHandlerFunctions) {
  handler := &SlackEventHandler{}
  handler.Handlers = handlers
  dp = handler
}

// Dispatch tries to resolve a valid handler. It writes and closes the request if possible, or returns a JSON encoded error if none found
func Dispatch(eventType string, params map[string]string, w http.ResponseWriter, r *http.Request) error {
  if dp == nil {
    return fmt.Errorf("Dispatcher not initialized")
  }

  if handler, ok := dp.Handlers[eventType]; ok {
    logging.Log.Debugf("Invoking handler for type: %s", eventType)
    handler.Function(params, w, r)
  } else {
    return fmt.Errorf("Unable to find handler for even type %s", eventType)
  }

  return nil
}
