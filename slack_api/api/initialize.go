package api


import (
  apiTypes "github.com/nakkiservo/slackwiener_backend/slack_api/types"
)

type SlackApiConfig struct {
  Authorization string // This will be our authorization token
}

// There must exist a more sensible way to do this, but i'm just doing this because it's friday
type SlackFilesAPIWrapper struct {
  Info func(string) (apiTypes.SlackFileInfo, error)
}

type SlackAPI struct {
  Config            SlackApiConfig
  File              SlackFilesAPIWrapper
}

func Initialize(token string) *SlackAPI {
  api := &SlackAPI{
    Config: SlackApiConfig{ Authorization: token },
    File:   SlackFilesAPIWrapper{},
  }

  api.File = SlackFilesAPIWrapper{
    Info: func(file_id string) (apiTypes.SlackFileInfo, error) {
      return GetFileInfo(api.Config.Authorization, file_id)
    },
  }

  // Wrap files api callbacks

  return api
}




