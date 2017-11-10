
package default_handlers

import (
  slackApi "github.com/nakkiservo/slackwiener_backend/slack_api/api/types"
  "github.com/nakkiservo/slackwiener_backend/logging"
  "net/http"
  "encoding/json"
)

func DefaultFileCreatedHandler(event slackApi.SlackEvent, w http.ResponseWriter, r *http.Request) {
  logging.Log.Debug("A file was created", event)
  w.WriteHeader(http.StatusOK)
  w.Header().Add("Content-Type", "application/json")
  json.NewEncoder(w).Encode("alles gut")
}

func DefaultFileSharedHandler(event slackApi.SlackEvent, w http.ResponseWriter, r *http.Request) {
  logging.Log.Debug("A file was shared!", event)
  w.WriteHeader(http.StatusOK)
  w.Header().Add("Content-Type", "application/json")
  json.NewEncoder(w).Encode("alles gut")
}
