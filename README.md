# imgur

Post images to imgur from go.

## Usage

```go
package main

import (
  "fmt"
  "os"
  "github.com/fiskhandlarn/imgur"
)

func (a *App) DeleteImage(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  hash := vars["hash"]

  // with bearer
  var bearer = "SUPERSECRET BEARER"
  err := imgur.Delete(data, &bearer)
  if err != nil {
    fmt.Fprintln(os.Stderr, "Deletion from imgur failed: " + err.Error())
    return
  }

  // respond with success
}

func (a *App) UploadImage(w http.ResponseWriter, r *http.Request) {
  file, _, err := r.FormFile("image")
  if err != nil {
    fmt.Fprintln(os.Stderr, err.Error())
    return
  }
  defer file.Close()

  data, err := ioutil.ReadAll(file)
  if err != nil {
    fmt.Fprintln(os.Stderr, "image error: " + err.Error())
    return
  }

  // upload anonymously
  imageURL, err := imgur.Upload(data, nil)

  // ... or with bearer
  var bearer = "SUPERSECRET BEARER"
  imageURL, err := imgur.Upload(data, &bearer)

  if err != nil {
    fmt.Fprintln(os.Stderr, "Upload to imgur failed: " + err.Error())
    return
  }

  // do something with imageURL
}

```
