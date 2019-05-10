package main

import (
   "fmt"
   "os"
   "net/http"
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
   fmt.Println("heres the hook Id", hook.Id)
   if err != nil {
      log.Fatal("fatal parsing hook", err)
      return
   }
   //pull latest in parent repository
   //run glue script in background
   //  redirect logs to some buffer so we can display
   //  the log through http 
   //  copy build artifacts to a directory within this repo.
}

func buildsHandler(w http.ResponseWriter, r *http.Request) {
   http.ServeFile(w, r, "./builds.html")
   fmt.Println("served builds page")
}

func main() {
   if os.Args[1] == "" {
      log.Fatal("Require arg - path to build script!")
      return
   }
   http.HandleFunc("/", homePage) // set router
   http.HandleFunc("/bootstrap", BootstrapPage) // set router
   http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
   http.HandleFunc("/gitwebhook", githubHandler) // set router
   http.HandleFunc("/builds", buildsHandler) // set router

   err := http.ListenAndServe(":80", nil) // set listen port
   if err != nil {
      log.Fatal("ListenAndServe: ", err)
   }
}
