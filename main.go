package main

import (
    "os"
    "log"
    "regexp"
    "net/http"
    "html/template"
)

func main() {
    bind := ":80"
    if len(os.Args) > 1 {
        bind = os.Args[1]
    }

    redirTmpl, err := template.ParseFiles("redirect.html")
    if err != nil { log.Fatal(err) }

    urlRegexp := regexp.MustCompile(`^/(\*?)((?:(?:[a-zA-Z0-9-]+\.)+[a-zA-Z0-9-]{2,}|(?:[0-9]+\.){3}[0-9]+)(?::[0-9]+)?(?:/.*)?)$`)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        m := urlRegexp.FindStringSubmatch(r.RequestURI)
        if m != nil {
            var url = "http://"
            if m[1] == "*" { url = "https://" }
            url += m[2]
            redirTmpl.Execute(w, url)
        } else if r.URL.Path == "/" {
            http.ServeFile(w, r, "index.html")
        } else {
            http.NotFound(w, r)
        }
    })
    http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "robots.txt")
    })
    http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "favicon.ico")
    })
    http.HandleFunc("/logo.png", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "logo.png")
    })
    http.HandleFunc("/script", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "script.js")
    })

    log.Fatal(http.ListenAndServe(bind, nil))
}
