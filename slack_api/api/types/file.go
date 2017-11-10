package types


// Wraps the slack files.info response to reponse status and file object
type FileResponseWrapper {
  ApiRequestResponse
  SlackFileInfo
}

// Slack file.info object, refer to the documentation
type SlackFileInfo struct {
  Id                  string  `json:"id"`
  Timestamp           int     `json:"timestamp"`
  Name                string  `json:"name"`
  Title               string  `json:"title"`
  Mimetype            string  `json:"mimetype"`
  Filetype            string  `json:"filetype"`
  PrettyType          string  `json:"pretty_type"`
  User                string  `json:"user"`
  Mode                string  `json:"mode"`
  Editable            bool    `json:"editable"`
  IsExternal          bool    `json:"is_external"`
  ExternalType        string  `json:"external_type"`
  Size                int     `json:"size"`
  Url                 string  `json:"url"`
  UrlDownload         string  `json:"url_download"`
  UrlPrivate          string  `json:"url_private"`
  UrlPrivateDownload  string  `json:"url_private_download"`
  IsPublic            bool    `json:"is_public"`
  PublicURLShared     bool    `json:"public_url_shared"`
}
