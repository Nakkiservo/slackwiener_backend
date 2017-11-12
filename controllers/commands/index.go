package commands

// Instead of building yet another dispatcher system or refactoring the old one
// We just handle some commands locally, if that's ok. It's ok, isn't it? Yeah it is
import (
  "net/http"
  "github.com/nakkiservo/slackwiener_backend/logging"
  "github.com/nakkiservo/slackwiener_backend/slack_api/commands"
)


// Index greps the command name and dispatches things
func Index(params map[string]string, w http.ResponseWriter, r *http.Request) {
  logging.Log.Debug("Action at commands controller index: ", params)

  w.WriteHeader(http.StatusOK)

  cmd, err := commands.DecodeCommandPayload(r)

  if err != nil {
    logging.Log.Errorf("Unable to decode payload: %s", err.Error())
    return
  }

  logging.Log.Debug("Decoded command payload: ", cmd)
}


