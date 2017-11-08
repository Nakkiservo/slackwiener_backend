package events


type SlackEvent struct {
  Token           string    `json:"token"`
  TeamID          string    `json:"team_id"`
  ApiAppID        string    `json:"api_app_id"`
  Type            string    `json:"type"`
  AuthedUsers     []string  `json:"authed_users"`
  EventID         string    `json:"event_id"`
  EventTime       int       `json:"event_time"`
  Challenge       string    `json:"challenge"` // only here for the handshake
  Event           map[string]interface{} `json:"event"`
}
