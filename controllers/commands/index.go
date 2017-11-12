package commands

// Instead of building yet another dispatcher system or refactoring the old one
// We just handle some commands locally, if that's ok. It's ok, isn't it? Yeah it is
import (
  "net/http"
  "github.com/nakkiservo/slackwiener_backend/logging"
  "github.com/nakkiservo/slackwiener_backend/slack_api/commands"
  "encoding/json"
)


// Index greps the command name and dispatches things
func Index(params map[string]string, w http.ResponseWriter, r *http.Request) {
  logging.Log.Debug("Action at commands controller index: ", params)


  cmd, err := commands.DecodeCommandPayload(r)

  if err != nil {
    logging.Log.Errorf("Unable to decode payload: %s", err.Error())
    w.WriteHeader(http.StatusOK)
    return
  }

  switch(cmd.Command) {
    case "/testcommand":
      HandleTestCommand(cmd, w)
    case "/testinchannel":
      HandleTestInChannelCommand(cmd,w)
    default:
      logging.Log.Debugf("Token: %s, Command: %s, Text: %s", cmd.Token, cmd.Command, cmd.Text)
  }
}

func HandleTestCommand(cmd *commands.SlackCommandPayload, w http.ResponseWriter) {

  w.Header().Set("Content-Type", "application/json")

  payload := make(map[string]interface{})

  payload["text"] = "You have invoked the TESTCOMMAND! OH NOES!"
  payload["attachments"] = []map[string]string{
    {
     "text": "Your arguments were: " + cmd.Text,
    },
  }

  json.NewEncoder(w).Encode(&payload)
}

func HandleTestInChannelCommand(cmd *commands.SlackCommandPayload, w http.ResponseWriter) {

  w.Header().Set("Content-Type", "application/json")

  payload := make(map[string]interface{})

  payload["text"] = cmd.UserName + " has invoked this spammy test command. We are truly sorry"
  payload["response_type"] = "in_channel"

  var payloadText string

  if cmd.Text == "" {
      payloadText = "The bastard gave no arguments either. :("
  } else {
      payloadText = "They thought the following arguments would do something: " + cmd.Text
  }

  payload["attachments"] = []map[string]string{
    {
     "text": payloadText,
    },
  }

  json.NewEncoder(w).Encode(&payload)
}


