package types

type ApiRequestResponse struct {
  Ok      bool              `json:"ok"`
  Error   interface{}       `json:"error"`
  Warning interface{}       `json:"warning`
}


