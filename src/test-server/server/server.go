package main

import (
  "fmt"
  "net/http"
  "test-server/router"
)

func main() {
    http.HandleFunc("/", router.RouteRequest)
    port := ":8080"

    fmt.Printf("Server listening on http://localhost%s \n", port)
    http.ListenAndServe(port, nil)
}
