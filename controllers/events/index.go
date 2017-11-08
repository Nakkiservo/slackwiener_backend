package events

import (
  //  "bitbucket.org/ncolabs/slackwiener_backend/config" :: todo: include for comparing with the slack token
  "bitbucket.org/ncolabs/slackwiener_backend/logging"
  "bitbucket.org/ncolabs/slackwiener_backend/utils"
  "bitbucket.org/ncolabs/slackwiener_backend/slack_events"
  "encoding/json"
  "net/http"
)

type SlackChallenge struct {
  Challenge   string `json:"challenge"`
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
//  conf := config.GetAppConfiguration()
  //slackToken :=  conf.SlackToken :: todo: compare token to request token
  var event *slack_events.SlackEvent


  if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
    logging.Log.Errorf("Unable to decode Slack request header: %s", err.Error())
     utils.SendError(w, utils.HttpError{
      Code:     http.StatusBadRequest,
      Message:  "Could not read request body: " + err.Error(),
    })
    return
  }

  logging.Log.Debug("Aww yess, we have a slack event", event)

  if event.Type == "url_verification" {
    response := &SlackChallenge{Challenge: event.Challenge}
    w.WriteHeader(http.StatusOK)
    w.Header().Add("Content-Type", "application/json")
    json.NewEncoder(w).Encode(&response)
  } else {
    logging.Log.Debug("Other event type: " ,event)
    w.WriteHeader(http.StatusOK)
  }

}

