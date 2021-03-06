package default_handlers

import (
  slackApi "github.com/nakkiservo/slackwiener_backend/slack_api/api"
  slackTypes "github.com/nakkiservo/slackwiener_backend/slack_api/types"
  "github.com/nakkiservo/slackwiener_backend/logging"
  "net/http"
  "encoding/json"
)

func DefaultMessageHandler(api *slackApi.SlackAPI, event slackTypes.SlackEvent, w http.ResponseWriter, r *http.Request) {
  logging.Log.Debug("Got a slack message event!", event)
  w.WriteHeader(http.StatusOK)
  w.Header().Add("Content-Type", "application/json")
  json.NewEncoder(w).Encode("alles gut")
}
