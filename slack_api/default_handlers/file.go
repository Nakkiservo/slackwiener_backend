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
// And by default we make the file public and store it's fields into our database
func DefaultFileSharedHandler(api *slackApi.SlackAPI, event slackTypes.SlackEvent, w http.ResponseWriter, r *http.Request) {
  logging.Log.Debug("A file was shared!", event)
  w.WriteHeader(http.StatusOK)
  w.Header().Add("Content-Type", "application/json")
  json.NewEncoder(w).Encode("alles gut")

  logging.Log.Debug("Trying to get file info")

  info, err := api.File.GetPublicLink(event.Event["file_id"].(string))
  if err != nil {
    logging.Log.Error("Unable to get file info")
  } else {
    logging.Log.Infof("Creating a public link entry for file %s", event.Event["file_id"].(string), info)
    logging.Log.Debug(event.Event)
    /*
    entry := db.SlackFileLink{
      FileId: event.Event["file_id"].(string),
      PrivateURL: info.Permalink,
      PublicURL: info.PermalinkPublic,
    }
    */

    /* Not needed since we don't really need slackwiener to get any info anymore, due to the fact that just setting the link public works just as well
    if dbErr := db.CreateLink(&entry); err != nil {
      logging.Log.Debugf("Unable to create link entry: %s", dbErr.Error())
    } else {
      logging.Log.Infof("Created a public link entry %s -> %s", entry.PrivateURL, entry.PublicURL)
    }
    */
  }
}
