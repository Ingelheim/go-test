package handler

import (
  "net/http"
)

type Handler interface {
    HandleRequest(w http.ResponseWriter, r *http.Request)
}
