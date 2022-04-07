package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":1923", nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to aiggato!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "For further questions, please send an email to <a href=\"mailto:alex.badragan@protonmail.com\">alex.badragan@protonmail.com</a>.")
	}

}
