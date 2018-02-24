package commands

// Instead of building yet another dispatcher system or refactoring the old one
// We just handle some commands locally, if that's ok. It's ok, isn't it? Yeah it is
import (
  "net/http"
  "github.com/nakkiservo/slackwiener_backend/logging"
  "github.com/nakkiservo/slackwiener_backend/slack_api/commands"
  "encoding/json"
  "time"
  "math/rand"
  "fmt"
  "regexp"
  "strconv"
)

var source *rand.Rand

// Index greps the command name and dispatches things
func Index(params map[string]string, w http.ResponseWriter, r *http.Request) {
  logging.Log.Debug("Action at commands controller index: ", params)

  cmd, err := commands.DecodeCommandPayload(r)

  if err != nil {
    logging.Log.Errorf("Unable to decode payload: %s", err.Error())
    w.WriteHeader(http.StatusOK)
    return
  }

  switch(cmd.Command) {
  case "/testcommand":
    HandleTestCommand(cmd, w)
  case "/testinchannel":
    HandleTestInChannelCommand(cmd,w)
  case "/noppa":
    HandleDiceCommand(cmd, w)
  case "/darra":
    HandleDarra(cmd, w)
  default:
    logging.Log.Debugf("Token: %s, Command: %s, Text: %s", cmd.Token, cmd.Command, cmd.Text)
  }
}

func HandleTestCommand(cmd *commands.SlackCommandPayload, w http.ResponseWriter) {

  w.Header().Set("Content-Type", "application/json")

  payload := make(map[string]interface{})

  payload["text"] = "You have invoked the TESTCOMMAND! OH NOES!"
  payload["attachments"] = []map[string]string{
    {
      "text": "Your arguments were: " + cmd.Text,
    },
  }

  json.NewEncoder(w).Encode(&payload)
}

func HandleTestInChannelCommand(cmd *commands.SlackCommandPayload, w http.ResponseWriter) {

  w.Header().Set("Content-Type", "application/json")

  payload := make(map[string]interface{})

  payload["text"] = cmd.UserName + " has invoked this spammy test command. We are truly sorry"
  payload["response_type"] = "in_channel"

  var payloadText string

  if cmd.Text == "" {
    payloadText = "The bastard gave no arguments either. :("
  } else {
    payloadText = "They thought the following arguments would do something: " + cmd.Text
  }

  payload["attachments"] = []map[string]string{
    {
      "text": payloadText,
    },
  }

  json.NewEncoder(w).Encode(&payload)
}

func HandleDiceCommand(cmd *commands.SlackCommandPayload, w http.ResponseWriter) {
  if source == nil {
    source = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
  }


  w.Header().Set("Content-Type", "application/json")
  payload := make(map[string]interface{})

  dRe := regexp.MustCompile("(\\d+)?d(\\d+)")


  result := RollDice(cmd.Text)

  var payloadText string

  if dRe.MatchString(cmd.Text) {
    payload["text"] = cmd.UserName + " rolls " + cmd.Text + "."
    payloadText = fmt.Sprintf("The completely random and fair result is: %d", result)
    payload["response_type"] = "in_channel"
  } else {
    payloadText = "Invalid dice roll"
  }

  payload["attachments"] = []map[string]string{
    {
      "text": payloadText,
    },
  }

  json.NewEncoder(w).Encode(&payload)

}


func HandleDarra(cmd *commands.SlackCommandPayload, w http.ResponseWriter) {

  w.Header().Set("Content-Type", "application/json");

  payload := make(map[string]interface{})
  payload["text"] = cmd.UserName + " has a head-splitting darra"
  payload["response_type"] = "in_channel"

  if cmd.Text != "" {
    var payloadText string

    if cmd.Text == "sherry" {
      payloadText = "It is the fearsome SHERRYDARRA";

      payload["attachments"] = []map[string]string{
        {
          "text": payloadText,
        },
      }
    }
  }

  json.NewEncoder(w).Encode(&payload)

}

func RollDice(dice string) int {
  dRe := regexp.MustCompile("(\\d+)?d(\\d+)")

  // 
  if dice != "" && dRe.MatchString(dice) {
    m := dRe.FindStringSubmatch(dice)

    times := 1
    var d int

    if m[1] != "" {
      tmp, err := strconv.Atoi(m[1])

      if err == nil {
        times = tmp
      }

    }

    d, err := strconv.Atoi(m[2])

    if err != nil {
      d = 6
    }

    if d <= 1 {
      d = 2
    }

    sum := 0

    for i := 0; i < times; i++ {
      sum += 1 + source.Intn(d - 1)
    }

    return sum
  }

  // we default to 1d6
  return 1 + source.Intn(5)
}

