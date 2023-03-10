package main

import (
  "log"
  "net/http"
  "time"
)
func delayedSampleJs(w http.ResponseWriter, req *http.Request) {
  log.Println("Request received for /block.js")
  log.Println("Sleep started at:", time.Now().String())
  for i := 1; i <= 350; i++ {
    time.Sleep(1 * time.Second)
    log.Println("Sleeping", i)
    w.Write([]byte(" "))
  }
  log.Println("Sleep ended at:", time.Now().String())
  req.Header.Set("Content-type", "application/javascript")
  http.ServeFile(w, req, "./block.js")
}

func callDelay(w http.ResponseWriter, req *http.Request) {
  log.Println("Request received for /test")
  log.Println("Sleep started at:", time.Now().String())
  for i := 1; i <= 400; i++ {
    time.Sleep(1 * time.Second)
    log.Println("Sleeping", i)
    w.Write([]byte(" "))
  }
  log.Println("Sleep ended at:", time.Now().String())
  req.Header.Set("Content-type", "application/javascript")
  http.ServeFile(w, req, "./index.html")
}

func main() {
  http.HandleFunc("/block.js", delayedSampleJs)
  // http.HandleFunc("/test", callDelay)
  http.Handle("/", http.FileServer(http.Dir("./")))
  log.Println("Starting server at port :9000...")
  http.ListenAndServe(":9000", nil)
}