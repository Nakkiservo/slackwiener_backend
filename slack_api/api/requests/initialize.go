package requests


type SlackApiConfig struct {
  Authorization string // This will be our authorization token
}


var config *SlackApiConfig

func Initialize(token string) {
  config = &SlackApiConfig{ Authorization: token }
}
