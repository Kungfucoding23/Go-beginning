package main

// servidor
import (
	"net/http"
)

func main() {

	// routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactPage)

	// start server
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contact Page"))
}
