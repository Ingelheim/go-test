package handler

import (
  "net/http"
)

type FaviconHandler struct {}

func (h FaviconHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "favicon.ico")
}
