package handler

import (
  "net/http"
  "html/template"
)

type HomeHandler struct {}

func (h HomeHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("views/home.html")
  t.Execute(w, h)
}
