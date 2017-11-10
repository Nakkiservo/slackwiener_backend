package main

import (
  "net/http"
  "bitbucket.org/ncolabs/slackwiener_backend/logging"
  "bitbucket.org/ncolabs/slackwiener_backend/config"
  "bitbucket.org/ncolabs/slackwiener_backend/routing"
  "bitbucket.org/ncolabs/slackwiener_backend/events"
  "github.com/urfave/negroni"
  "strconv"
)

func main() {
  logging.Initialize();

  logging.Log.Info("Starting SlackWiener backend server")
  logging.Log.Info("Loading configuration")

  conf := config.GetAppConfiguration()


  logging.Log.Debug("Configuring routes")

  r := routing.CreateRouter(routing.Routes, routing.Handlers)

  events.InitializeDispatcher(event.DefaultHandlers)

  n := negroni.Classic() // Includes some default middlewares
  n.UseHandler(r)

  server := &http.Server{
    Handler: n,
    Addr: conf.ListenAddress + ":" + strconv.Itoa(conf.ListenPort),
  }

  logging.Log.Info("Starting server")
  if err := server.ListenAndServe(); err != nil {
    logging.Log.Criticalf("Unable to start HTTP server: %s", err.Error())
    panic("UNABLE TO START SERVER!")
  }

  logging.Log.Infof("Started server at port %d", conf.ListenPort)
}
