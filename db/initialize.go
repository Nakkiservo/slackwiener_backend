package db

import (
  "github.com/nakkiservo/slackwiener_backend/logging"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  "time"
)

// db is the main database connection
var db *gorm.DB

// A slack public file link
type SlackFileLink struct {
  ID          uint  `gorm:"primary_key" json:"id"`
  CreatedAt   time.Time `json:"created_at"`
  UpdatedAt   time.Time `json:"updated_at"`
  FileId      string    `json:"file_id"`
  PublicURL   string    `json:"public_url"`
  PrivateURL  string    `json:"private_url"`
}

// Initialize initializes the initialization of initial database. At least initially
func Initialize() {
  var err error
  db, err = gorm.Open("sqlite3", "./slackwiener.db")
  if err != nil {
    logging.Log.Criticalf("Unable to open database connection: %s", err)
    panic("UNABLE TO OPEN DATABASE!")
  }

  db.LogMode(true)

  if !db.HasTable(&SlackFileLink{}) {
    logging.Log.Info("Creating User model table")
    db.CreateTable(&SlackFileLink{})
  }
}


// CreateLink creates a new link entry in the database
func CreateLink(link *SlackFileLink) error {
  err := db.Create(link).Error
  return err
}


