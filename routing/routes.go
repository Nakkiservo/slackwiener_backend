package routing

import (
  "github.com/nakkiservo/slackwiener_backend/controllers/events"
  "github.com/nakkiservo/slackwiener_backend/controllers/commands"
)


var Handlers HandlerFunctions = HandlerFunctions{
  "incomingEvents":      {Function: events.Events  },
  "incomingCommands":    {Function: commands.Index },
}

var Routes HandlerRoutes = HandlerRoutes{
  "POST /events/": "incomingEvents",
  "POST /commands/{cmd:\\w+}": "incomingCommands",
}

