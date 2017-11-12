// package Commands adds Slack slash-command payload parsing and required types
package commands


import (
  "net/http"
  "encoding/json"
  "fmt"
)

// From slack api examples:
// token=gIkuvaNzQIHg97ATvDxqgjtO
// team_id=T0001
// team_domain=example
// enterprise_id=E0001
// enterprise_name=Globular%20Construct%20Inc
// channel_id=C2147483705
// channel_name=test
// user_id=U2147483697
// user_name=Steve
// command=/weather
// text=94070
// response_url=https://hooks.slack.com/commands/1234/5678
// trigger_id=13345224609.738474920.8088930838d88f008e0

var _payloadKeys []string = []string{
  "token",
  "team_id",
  "team_domain",
  "enterprise_id",
  "enterprise_name",
  "channel_id",
  "channel_name",
  "user_id",
  "user_name",
  "command",
  "text",
  "response_url",
}

type SlackCommandPayload struct {
  Token           string `json:"token"`
  TeamID          string `json:"team_id"`
  TeamDomain      string `json:"team_domain"`
  EnterpriseID    string `json:"enterprise_id"`
  EnterpriseName  string `json:"enterprise_name"`
  ChannelID       string `json:"channel_id"`
  ChannelName     string `json:"channel_name"`
  UserID          string `json:"user_id"`
  UserName        string `json:"user_name"`
  Command         string `json:"command"`
  Text            string `json:"text"`
  ResponseURL     string `json:"reponse_url"`
  TriggerID       string `json:"trigger_id"`
}

// DecodeCommandPayload decodes a http.Request form and tries to return a SlackCommandPayload
func DecodeCommandPayload(r *http.Request) (*SlackCommandPayload, error) {

  if err := r.ParseForm(); err != nil {
    return nil, fmt.Errorf("Unable to parse command payload: %s", err.Error())
  }

  values := map[string]string{}

  for _, key := range _payloadKeys {
    values[key] = r.PostFormValue(key)
  }

  rawJson,_ := json.Marshal(values)
  payload   := SlackCommandPayload{}

  if jsErr := json.Unmarshal(rawJson, &payload); jsErr != nil {
    return nil, fmt.Errorf("Error while parsing form data: %s", jsErr.Error())
  }

  return &payload, nil
}
