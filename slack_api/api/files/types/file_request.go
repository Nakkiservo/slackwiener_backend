package types


type FileResponse {
  Ok      bool          `json:"ok"`
  File    FileStruct    `json:"file"`
  Error   interface{}   `json:"error"`
}

type FileStruct struct {
  id            string
  timestamp     int
  name          string
  title         string
  mimetype      string
  filetype      string
  pretty_type   string
  user          string
  mode          string
  editable      bool
  is_external   bool
  external_type string
  size          int
  url           string
  url_download  string
  url_private   string
  url_private_download string
  is_public     bool
  public_url_shared bool
}
