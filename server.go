package main

import (
	"fmt"
	"net/http"
)

func bowerHandler(w http.ResponseWriter, r *http.Request) {
	path := fmt.Sprintf("bower_components/%v", r.URL.Path[1:])
	fmt.Println(path)
	http.ServeFile(w, r, path)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.Handle("/bower_components/", http.StripPrefix("/bower_components/", http.FileServer(http.Dir("bower_components"))))
	http.Handle("/my-components/", http.StripPrefix("/my-components/", http.FileServer(http.Dir("my-components"))))
	// http.HandleFunc("/bower_components/.+", bowerHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// func main() {
// 	// Simple static webserver:
// 	http.ListenAndServe(":8080", http.FileServer(http.Dir("index.html")))
// }
