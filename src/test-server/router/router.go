package router

import (
  "fmt"
  "net/http"
  "test-server/handler"
)

var handlers = map[string]handler.Handler{
  "/": &handler.HomeHandler{},
  "fallback": &handler.FallbackHandler{},
}

func handleFallbackRequest(w http.ResponseWriter, r *http.Request) {
  handlers["fallback"].HandleRequest(w, r)
}

func RouteRequest(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path

    fmt.Printf("Handling request for path %s \n", path)

    if handler, ok := handlers[path]; ok {
      handler.HandleRequest(w, r)
    } else {
      handleFallbackRequest(w, r)
    }
}
