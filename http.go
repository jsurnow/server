package main

import (
    "net/http"
    "io"
)

func hello (res http.ResponseWriter, req *http.Request) {
    res.Header().Set(
        "Content-Type",
        "application/json",
    )
    io.WriteString(res, "{hello:world}\n")
}

func main() {
    http.HandleFunc("/hello", hello);
    http.ListenAndServe(":9000", nil);
}
