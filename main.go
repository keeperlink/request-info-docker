package main

import (
    "fmt"
    "net/http"
    "sort"
)

var version = "dev"

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Server", "request-info-docker:"+version)
    w.Header().Set("X-App-Version", version)
    fmt.Fprintln(w, "<html><body>");

    fmt.Fprintln(w, "<i>Request Host</i>: ", r.Host, "</br>")
    fmt.Fprintln(w, "<i>Request URL</i>: ", r.URL, "</br>")
    fmt.Fprintln(w, "<i>Context path</i>: ", r.URL.Path[1:], "</br>")
    fmt.Fprintln(w, "<i>Client IP (RemoteAddr)</i>: ", r.RemoteAddr, "</br>")
    fmt.Fprintln(w, "<i>Request TLS</i>: ", r.TLS, "</br>")

    fmt.Fprintln(w, "</br><b>Request Headers:</b></br>")
    var keys []string
    for k := range r.Header {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    fmt.Fprintln(w, "<ul>")
    for _, k := range keys {
        fmt.Fprintln(w, "<li><i>", k, "</i>:", r.Header[k], "</li>")
    }
    fmt.Fprintln(w, "</ul>")

    fmt.Fprintln(w, "</br><b>Form parameters:</b></br>")
    r.ParseForm()
    keys = []string{}
    for k := range r.Form {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    fmt.Fprintln(w, "<ul>")
    for _, k := range keys {
        fmt.Fprintln(w, "<li><i>", k, "</i>:", r.Form[k], "</li>")
    }
    fmt.Fprintln(w, "</ul>")

    fmt.Fprintln(w, "</body></html>");
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
