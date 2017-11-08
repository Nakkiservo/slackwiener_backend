package routing

import (
  "bitbucket.org/ncolabs/slackwiener_backend/controllers/events"
)

var Handlers HandlerFunctions = HandlerFunctions{
  "incomingEvents":      {Function: events.Events  },
}

var Routes HandlerRoutes = HandlerRoutes{
  "POST /events/": "incomingEvents",
}

