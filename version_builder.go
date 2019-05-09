package main

import (
   "fmt"
   "net/http"
   //"net/http/httputil"
   "log"
)

func homePage(w http.ResponseWriter, r *http.Request) {
   http.ServeFile(w, r, "./index.html")
   fmt.Println("served home page")
}

func BootstrapPage(w http.ResponseWriter, r *http.Request) {
   http.ServeFile(w, r, "./bootstrap.html")
   fmt.Println("served bootstrap page")
}

func main() {
   http.HandleFunc("/", homePage) // set router
   http.HandleFunc("/bootstrap", BootstrapPage) // set router
   http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
   err := http.ListenAndServe(":8090", nil) // set listen port
   if err != nil {
      log.Fatal("ListenAndServe: ", err)
   }
}
