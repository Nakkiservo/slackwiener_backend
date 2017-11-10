package events

import (
  //  "bitbucket.org/ncolabs/slackwiener_backend/config" :: todo: include for comparing with the slack token
  "bitbucket.org/ncolabs/slackwiener_backend/logging"
  "bitbucket.org/ncolabs/slackwiener_backend/utils"
  slackDispatcher "bitbucket.org/ncolabs/slackwiener_backend/slack_api/dispatcher"
  slackApi "bitbucket.org/ncolabs/slackwiener_backend/slack_api/api/types"
  "encoding/json"
  "net/http"
)

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
  var event *slackApi.SlackEvent


  if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
    logging.Log.Errorf("Unable to decode Slack request header: %s", err.Error())
    utils.SendError(w, utils.HttpError{
      Code:     http.StatusBadRequest,
      Message:  "Could not read request body: " + err.Error(),
    })
    return
  }

  if event.Type == "url_verification" {
    response := &slackApi.SlackChallenge{Challenge: event.Challenge}
    w.WriteHeader(http.StatusOK)
    w.Header().Add("Content-Type", "application/json")
    json.NewEncoder(w).Encode(&response)
  } else if event.Type == "event_callback" {
    if dispatchError := slackDispatcher.Dispatch(event.Event["type"].(string), *event,w,r); dispatchError != nil {
      logging.Log.Errorf("Error handling event type '%s': %s", event.Type, dispatchError.Error())
      utils.SendError(w, utils.HttpError{
        Code:     http.StatusBadRequest,
        Message:  "Error handling request: " + dispatchError.Error(),
      })
      return
    }

    // We can safely just return here since the handlers call their own closers based on the content they want to send
  }

}

