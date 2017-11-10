package events

import (
  dh "bitbucket.org/ncolabs/slackwiener_backend/slack_api/default_handlers"
)

var DefaultHandlers SlackHandlerFunctions = SlackHandlerFunctions {
  "message" { Function: dh.DefaultMessageHandler },
}
