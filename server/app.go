package main

import (
    "github.com/kisom/gopush_git/pushover"
    "fmt"
    "log"
    "net/http"
    "os"
)

func notify(s string) bool {
    identity := pushover.Authenticate(os.Getenv("PO_API"),
                                      os.Getenv("PO_USER"))
    if len(s) == 0 {
           s = "heroku visitor trigger"
    }
    return pushover.Notify(identity,  s)
}

func main() {
    http.HandleFunc("/hello", hello_ip)
    http.HandleFunc("/", hello)
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "hello, world!")
}

func hello_ip(w http.ResponseWriter, req *http.Request) {
    if notify(fmt.Sprintf("heroku: hello from %s", req.RemoteAddr)) {
            fmt.Fprintln(w, "hello sent!")
    } else {
            fmt.Fprintln(w, "something bad has happened!")
    }
}
