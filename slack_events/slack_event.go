package slack_events

import (
  "encoding/json"
)

type SlackEvent struct {
  Token           string    `json:"token"`
  TeamID          string    `json:"team_id"`
  ApiAppID        string    `json:"api_app_id"`
  Type            string    `json:"type"`
  AuthedUsers     []string  `json:"authed_users"`
  EventID         string    `json:"event_id"`
  EventTime       int       `json:"event_time"`
  Event           interface{} `json:"-"`
}
