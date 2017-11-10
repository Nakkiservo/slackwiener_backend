// Package events provides an API similar to routing in order to connect different kinds of events their proper handlers
package dispatcher

import (
  "net/http"
  "fmt"
  slackApi "github.com/nakkiservo/slackwiener_backend/slack_api/types"
  "github.com/nakkiservo/slackwiener_backend/logging"
)

// SlackEventHandler describes a handler for a slack event. Function is the callback
type SlackEventHandlerFunc struct {
  Function      func(slackApi.SlackEvent, http.ResponseWriter, *http.Request)
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
func Dispatch(eventType string, event slackApi.SlackEvent, w http.ResponseWriter, r *http.Request) error {
  if dp == nil {
    return fmt.Errorf("Dispatcher not initialized")
  }

  if handler, ok := dp.Handlers[eventType]; ok {
    logging.Log.Debugf("Invoking handler for type: %s", eventType)
    handler.Function(event, w, r)
  } else {
    return fmt.Errorf("Unable to find handler for even type %s", eventType)
  }

  return nil
}

