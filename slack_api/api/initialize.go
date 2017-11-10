package api


type SlackApiConfig struct {
  Authorization string // This will be our authorization token
}

var config *SlackApiConfig

type SlackAPI struct {
  Config      SlackApiConfig
}

func Initialize(token string) SlackApi {
  api := &SlackAPI{
    Config: SlackApiConfig{
      Authorization: token
    }
  }
}



