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
    } else {
      if jErr := json.NewDecoder(resp.Body).Decode(&fileResp); jErr != nil {
        logging.Log.Debugf("Error parsing json: %s", jErr.Error())
        return slackTypes.SlackFileInfo{}, jErr
      } else {
        if fileResp.Ok {
          logging.Log.Debug("Got legit file information!", fileResp.File)
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
