package handler

import (
  "fmt"
  "net/http"
)

type FallbackHandler struct {}

func (h FallbackHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("Fallback\n")
  fmt.Fprintf(w, "404")
}
