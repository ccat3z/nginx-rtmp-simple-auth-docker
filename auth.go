package main

import (
    "log"
    "os"
    "net/http"
)

var pass string

func auth(w http.ResponseWriter, r *http.Request) {
    log.Println(r.URL)
    if (pass == r.URL.Query().Get("pass")) {
        w.WriteHeader(http.StatusOK)
    } else {
        w.WriteHeader(http.StatusForbidden)
    }
}

func main () {
    pass = os.Getenv("AUTH_PASS")

    if (pass == "") {
        log.Println("need AUTH_PASS environment")
        return
    }

    http.HandleFunc("/", auth)

    if err := http.ListenAndServe("localhost:8008", nil); err != nil {
        log.Fatal(err.Error())
    }
}
