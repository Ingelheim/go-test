package handler

import (
  "fmt"
  "net/http"
  "test-server/template"
)

type HomeHandler struct {}

func (h HomeHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("Hi route\n")
  var temp = &template.HomeTemplate{}
  temp.Title = "Lukas"
  temp.Render()
  fmt.Fprintf(w, "Hi")
}
