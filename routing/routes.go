package routing

import (
  "github.com/nakkiservo/slackwiener_backend/controllers/events"
)

var Handlers HandlerFunctions = HandlerFunctions{
  "incomingEvents":      {Function: events.Events  },
}

var Routes HandlerRoutes = HandlerRoutes{
  "POST /events/": "incomingEvents",
}

