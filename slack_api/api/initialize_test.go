package api_test


import (
  "testing"
  "github.com/nakkiservo/slackwiener_backend/slack_api/api"
  "fmt"
)

func TestCallback(t *testing.T) {
  testApi := api.Initialize("HOmouden tokeni!")

  fmt.Println(testApi)
}
