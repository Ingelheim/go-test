package handler

import (
  "net/http"
  "html/template"
)

type FallbackHandler struct {}

func (h FallbackHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("views/404.html")
  t.Execute(w, h)
}
