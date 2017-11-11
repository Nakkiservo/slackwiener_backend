package api

import (
  "github.com/nakkiservo/slackwiener_backend/logging"
  slackTypes "github.com/nakkiservo/slackwiener_backend/slack_api/types"
  "net/http"
  "net/url"
  "bytes"
  "strconv"
  "encoding/json"
  "fmt"
)

type FileJSONArguments struct {
  Token string `json:"token"` // App token
  File  string `json:"file"` // file_id acquired from backend
}

// GetFilePublicLink makes a file publicly readable and returns a SlackFileInfo struct 
// that has its permalink_public_url set. This can be used later
func GetFilePublicLink(token string, file_id string) (slackTypes.SlackFileInfo, error) {
  apiUrl := "https://slack.com/api/files.sharedPublicURL"
  data := FileJSONArguments{File: file_id}
  b := new (bytes.Buffer)
  json.NewEncoder(b).Encode(&data)

  r, _ := http.NewRequest("POST", apiUrl, b)
  r.Header.Add("Content-Type", "application/json")
  r.Header.Add("Authorization", "Bearer " + token)

  client := http.Client{}

  resp, err := client.Do(r)
  fileResp := slackTypes.FileResponseWrapper{}

  if err != nil {
    logging.Log.Errorf("API POST error: %s", err.Error())
    return slackTypes.SlackFileInfo{}, err
  } else {
    if resp.StatusCode != http.StatusOK {
      logging.Log.Errorf("HTTP status not ok: %d", resp.Status)
    } else {
      if jErr := json.NewDecoder(resp.Body).Decode(&fileResp); jErr != nil {
        logging.Log.Debugf("Error parsing json: %s", jErr.Error())
        return slackTypes.SlackFileInfo{}, jErr
      } else {
        if fileResp.Ok {
          logging.Log.Debug("Got file information", fileResp.File)
          return fileResp.File, nil
        } else {
          logging.Log.Error("Received an error from backend: %s", fileResp.Error)
          return slackTypes.SlackFileInfo{}, fmt.Errorf("Backend error: %s", fileResp.Error)
        }
      }
    }
  }

  return slackTypes.SlackFileInfo{}, nil

}

// GetFileInfo returns a populated SlackFileInfo struct regarding the provided file_id, or an error
func GetFileInfo(token string, file_id string) (slackTypes.SlackFileInfo, error) {
  apiUrl := "https://slack.com/api/files.info"

  data := url.Values{}
  data.Add("token", token)
  data.Add("file", file_id)

  client := http.Client{}

  r,_ := http.NewRequest("POST", apiUrl, bytes.NewBufferString(data.Encode()))
  r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

  resp, err := client.Do(r)
  fileResp := slackTypes.FileResponseWrapper{}

  if err != nil {
    logging.Log.Errorf("API POST error: %s", err.Error())
    return slackTypes.SlackFileInfo{}, err
  } else {
    if resp.StatusCode != http.StatusOK {
      logging.Log.Errorf("HTTP status not ok: %d", resp.Status)
      return slackTypes.SlackFileInfo{}, fmt.Errorf("HTTP status not ok: %d", resp.StatusCode)
    } else {
      if jErr := json.NewDecoder(resp.Body).Decode(&fileResp); jErr != nil {
        logging.Log.Debugf("Error parsing json: %s", jErr.Error())
        return slackTypes.SlackFileInfo{}, jErr
      } else {
        if fileResp.Ok {
          return fileResp.File, nil
        } else {
          logging.Log.Error("Received an error from backend: %s", fileResp.Error)
          return slackTypes.SlackFileInfo{}, fmt.Errorf("Backend error: %s", fileResp.Error)
        }
      }
    }
  }

  return slackTypes.SlackFileInfo{}, nil
}
