package dispatcher

import (
  dh "bitbucket.org/ncolabs/slackwiener_backend/slack_api/default_handlers"
)
/*
var Handlers HandlerFunctions = HandlerFunctions{
  "incomingEvents":      {Function: events.Events  },
}
*/

var DefaultHandlers SlackHandlerFunctions = SlackHandlerFunctions{
  "message": { Function: dh.DefaultMessageHandler },
}
