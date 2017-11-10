package default_handlers


import (
  apiTypes "bitbucket.org/ncolabs/slackwiener_backend/slack_api/api/types"
  "bitbucket.org/ncolabs/slackwiener_backend/logging"
  "net/http"
)

func Message(params map[string]string, w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Header().Add("Content-Type", "application/json")
  json.NewEncoder(w).Encode(&response)
}
