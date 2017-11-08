package index

import (
  "bitbucket.org/ncolabs/slackwiener_backend/config"
  "bitbucket.org/ncolabs/slackwiener_backend/logging"
  "bitbucket.org/ncolabs/slackwiener_backend/utils"
  se "bitbucket.org/ncolabs/slackwiener_backend/slack_events"
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
  conf := config.GetAppConfiguration()
  slackToken :=  conf.SlackToken
  var typeHeader EventTypeHeader


  if err := json.NewDecoder(&se.SlackEvent).Decode(&typeHeader); err != nil {
    logging.Log.Errorf("Unable to decode Slack request header: %s", err.Error())
     utils.SendError(w, utils.HttpError{
      Code:     http.StatusBadRequest,
      Message:  "Could not read request body: " + err.Error(),
    })
    return
  }

  if(typeHeader.type == 
  w.WriteHeader(http.StatusOK)
  w.Header().Add("Content-Type", "application/json")

}

