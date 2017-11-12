package main

import (
  "net/http"
  "github.com/nakkiservo/slackwiener_backend/logging"
  "github.com/nakkiservo/slackwiener_backend/config"
  "github.com/nakkiservo/slackwiener_backend/routing"
  "github.com/nakkiservo/slackwiener_backend/db"
  slackDispatcher "github.com/nakkiservo/slackwiener_backend/slack_api/dispatcher"
  slackAPI "github.com/nakkiservo/slackwiener_backend/slack_api/api"
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

  api := slackAPI.Initialize(conf.SlackToken)

  slackDispatcher.InitializeDispatcher(slackDispatcher.DefaultHandlers, api)

  db.Initialize() // Initialize db

  n := negroni.Classic() // Includes some default middlewares
  n.UseHandler(r)

  server := &http.Server{
    Handler: n,
    Addr: conf.ListenAddress + ":" + strconv.Itoa(conf.ListenPort),
  }

  logging.Log.Info("Starting server")
   if err := srv.ListenAndServeTLS("server.crt", "server.key"); err != nil {
    logging.Log.Criticalf("Error creating the server: %s", err.Error())
   }

  logging.Log.Infof("Started server at port %d", conf.ListenPort)
}
