package config

import (
  "github.com/BurntSushi/toml"
  "bitbucket.org/ncolabs/slackwiener_backend/logging"
  "sync"
)

// Only ever parse once
var once sync.Once
var config *AppConfig

type AppConfig struct {
  ListenAddress string
  ListenPort    int
  SlackToken    string
}


func GetAppConfiguration() *AppConfig {
  once.Do(loadTomlConfiguration)

  if config == nil {
    logging.Log.Critical("Unable to load configuration file ")
    panic("Failed to load configuration file!")
  }

  return config
}


func loadTomlConfiguration() {

  logging.Log.Debug("Loading configuration")

  var conf *AppConfig = &AppConfig{}

  if _, err := toml.DecodeFile("app_config.toml", conf); err != nil {
    logging.Log.Critical("Unable to load configuration file " + err.Error())
    panic("Failed to load configuration file! " + err.Error())
  }

  logging.Log.Debugf("Loaded configuration: %+v\n", *conf)
  config = conf
}



