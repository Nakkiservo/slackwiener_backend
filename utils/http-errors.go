package utils

import (
  "encoding/json"
  "net/http"
  "bitbucket.org/ncolabs/slackwiener_backend/logging"
  "fmt"
)

type HttpError struct {
  Code    int
  Message string
}


//SendError sends a JSON encoded error object along with appropriate status code
func SendError(w http.ResponseWriter, httpError HttpError) {
  SetContentTypeJSON(w)
  w.WriteHeader(httpError.Code)
  fmt.Fprint(w, httpError.toJson())
}

//SetContentTypeJSON sets the responsewriter's content type to json
func SetContentTypeJSON(w http.ResponseWriter) {
  w.Header().Set("Content-Type", "application/json")
}

//NotImplementedYet send a "not implemented message"
func NotImplementedYet(w http.ResponseWriter) {
  w.WriteHeader(http.StatusNotImplemented)
  fmt.Fprint(w, "Feature not implemented yet.\n")
}

func SendUnauthorized(w http.ResponseWriter) {
  w.WriteHeader(http.StatusUnauthorized)
}

func (h *HttpError) toJson() string {

  data, err := json.Marshal(h)

  if err != nil {
    return "{\"message\": \"Failed to marshal error message\"}"
  }

  return string(data)
}

func ControllerError(err error, w http.ResponseWriter, code int, message string) bool {
  if err != nil {
    logging.Log.Error("Controller error: " + err.Error())
    SendError(w, HttpError{
      Code: code,
      Message: message,
    })
    return true
  }

  return false
}
