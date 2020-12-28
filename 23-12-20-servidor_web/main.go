package main

import "net/http"

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil) //puerto 3000
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
