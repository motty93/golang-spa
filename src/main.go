package main

import (
	"app/src/handler"

	"net/http"
)

func main() {
	http.HandleFunc("/hello", handler.Hello)
	http.HandleFunc("/test/", handler.Static)
	http.Handle("/web/static/", http.StripPrefix("/web/static/", http.FileServer(http.Dir("web/static"))))

	http.ListenAndServe(":3000", nil)
}
