package main

import (
  "net/http"
  "github.com/nakkiservo/slackwiener_backend/logging"
  "github.com/nakkiservo/slackwiener_backend/config"
  "github.com/nakkiservo/slackwiener_backend/routing"
  slackDispatcher "github.com/nakkiservo/slackwiener_backend/slack_api/dispatcher"
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

  slackDispatcher.InitializeDispatcher(slackDispatcher.DefaultHandlers)

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
