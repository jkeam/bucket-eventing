package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

func main() {
  log.Print("Server starting at :8080")
  http.HandleFunc("/", server)
  http.ListenAndServe(":8080", nil)
}

func server(w http.ResponseWriter, r *http.Request) {
  bytedata, err := io.ReadAll(r.Body)
  reqBodyString := string(bytedata)
  if err != nil {
    log.Fatal(err)
    http.Error(w, err.Error(), 500)
    return
  }

  log.Println(reqBodyString)
  fmt.Fprintf(w, "%s", reqBodyString)
}
