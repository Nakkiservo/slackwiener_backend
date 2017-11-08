package index

import (
  "bitbucket.org/ncolabs/slackwiener_backend/config"
  "bitbucket.org/ncolabs/slackwiener_backend/logging"
  "encoding/json"
  "net/http"
)

// Used only to decode the received event type. TODO: move to somewere sensible
type EventTypeHeader struct {
  Type string `json:"type"`
}

func Index(params map[string]string, w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "views/index.html")
}


func Assets(params map[string]string, w http.ResponseWriter, r *http.Request) {
  handler := http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/")))
  handler.ServeHTTP(w,r)
}


// Events handles all the slack api events http callbacks
func Events(params map[string]string, w http.ResponseWriter, r *http.Request) {
  conf := config.GetAppConfiguration()
  slackToken :=  conf.SlackToken
  var typeHeader EventTypeHeader


  if err := json.NewDecoder(&EventTypeHeader).Decode(&typeHeader); err != nil {
    logging.Log.Errorf("Unable to decode Slack request header: %s", err.Error())
  }
}

