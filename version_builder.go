package main

import (
   "fmt"
   "net/http"
   //"net/http/httputil"
   "log"
   "gopkg.in/rjz/githubhook.v0"
)

func homePage(w http.ResponseWriter, r *http.Request) {
   http.ServeFile(w, r, "./index.html")
   fmt.Println("served home page")
}

func BootstrapPage(w http.ResponseWriter, r *http.Request) {
   http.ServeFile(w, r, "./bootstrap.html")
   fmt.Println("served bootstrap page")
}

func githubHandler(w http.ResponseWriter, r *http.Request) {
   fmt.Println("got the github hook!")
   secret := []byte("imasecret")
   hook, err := githubhook.Parse(secret, r)
   fmt.Println("heres the hook payload", hook.Payload)
   if err != nil {
      log.Fatal("fatal parsing hook", err)
   }
}

func main() {
   http.HandleFunc("/", homePage) // set router
   http.HandleFunc("/bootstrap", BootstrapPage) // set router
   http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
   http.HandleFunc("/gitwebhook", githubHandler) // set router

   err := http.ListenAndServe(":80", nil) // set listen port
   if err != nil {
      log.Fatal("ListenAndServe: ", err)
   }
}
