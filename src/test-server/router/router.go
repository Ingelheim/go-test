package router

import (
  "fmt"
  "net/http"
  "strings"
  "test-server/handler"
)

var handlers = map[string]handler.Handler{
  "/": &handler.HomeHandler{},
  "/favicon.ico": &handler.FaviconHandler{},
  "fallback": &handler.FallbackHandler{},
  // "public": &handler.PublicHandler{},
}

func handleFallbackRequest(w http.ResponseWriter, r *http.Request) {
  handlers["fallback"].HandleRequest(w, r)
}

func handlePublicRequest(w http.ResponseWriter, r *http.Request, filePath string) {
  handler.PublicHandler{}.HandleFileRequest(w, r, filePath)
}

func RouteRequest(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    public := "/public"

    fmt.Printf("Handling request for path %s \n", path)

    if strings.Contains(path, public) {
      subpath := path[len(public):]

      fmt.Printf("subpath %s \n", subpath)
      handlePublicRequest(w, r, path[1:])
    }

    if handler, ok := handlers[path]; ok {
      handler.HandleRequest(w, r)
    } else {
      handleFallbackRequest(w, r)
    }
}
