package handler

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body>hello world</body></html>")
}

func Static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/static/nojs.html")
}
