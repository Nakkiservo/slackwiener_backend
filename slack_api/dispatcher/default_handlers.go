package dispatcher

import (
  dh "github.com/nakkiservo/slackwiener_backend/slack_api/default_handlers"
)
/*
var Handlers HandlerFunctions = HandlerFunctions{
  "incomingEvents":      {Function: events.Events  },
}
*/

var DefaultHandlers SlackHandlerFunctions = SlackHandlerFunctions{
  "message": { Function: dh.DefaultMessageHandler },
  "file_created": { Function: dh.DefaultFileCreatedHandler},
  "file_shared": { Function: dh.DefaultFileSharedHandler},
}
