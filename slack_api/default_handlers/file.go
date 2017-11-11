
package default_handlers

import (
  slackTypes "github.com/nakkiservo/slackwiener_backend/slack_api/types"
  slackApi "github.com/nakkiservo/slackwiener_backend/slack_api/api"
  "github.com/nakkiservo/slackwiener_backend/logging"
  "net/http"
  "encoding/json"
)

// DefaultFileCreatedHandler is the default handler for file creation
func DefaultFileCreatedHandler(api *slackApi.SlackAPI, event slackTypes.SlackEvent, w http.ResponseWriter, r *http.Request) {
  logging.Log.Debug("A file was created", event)
  w.WriteHeader(http.StatusOK)
  w.Header().Add("Content-Type", "application/json")
  json.NewEncoder(w).Encode("alles gut")
}

// DefaultFileSharedHandler is the default handler for file sharing events
func DefaultFileSharedHandler(api *slackApi.SlackAPI, event slackTypes.SlackEvent, w http.ResponseWriter, r *http.Request) {
  logging.Log.Debug("A file was shared!", event)
  w.WriteHeader(http.StatusOK)
  w.Header().Add("Content-Type", "application/json")
  json.NewEncoder(w).Encode("alles gut")
}
