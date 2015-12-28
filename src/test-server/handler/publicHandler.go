package handler

import (
  "net/http"
  "fmt"
  "os"
  "time"
)

type PublicHandler struct {}

func (h PublicHandler) HandleFileRequest(w http.ResponseWriter, r *http.Request, filePath string) {
  fmt.Printf("Serving File %s \n", filePath)

  f, err := os.Open(filePath)
   if err != nil {
     fmt.Println(err)
   }
   defer f.Close()

   modTime := time.Now()

   if fi, err := os.Stat(filePath); err == nil {
     modTime = fi.ModTime()
   }

  http.ServeContent(w, r, filePath, modTime, f)
}
