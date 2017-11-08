package logging

import (
  "github.com/op/go-logging"
  "os"
)

var Log = logging.MustGetLogger("slackwiener_backend")

var format = logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`)



func Initialize() {

  backend := logging.NewLogBackend( os.Stderr, "", 0)
  backendFormatter := logging.NewBackendFormatter(backend, format)

  logging.SetBackend(backendFormatter)

}
