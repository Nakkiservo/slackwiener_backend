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

// Used to parse the inner even type. TODO: move somewhere sensible
type SlackInnerEventTypeWrapper struct {
  Type      string `json:"type"`
}

// We just wrap the challenge resonse here. TODO: move somewhere sensible
type SlackChallenge struct {
  Challenge   string `json:"challenge"`
}


