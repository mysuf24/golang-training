package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8080"

func main() {
	http.HandleFunc("/", router)

	fmt.Println("Server berjalan di http://localhost" + PORT)
	http.ListenAndServe(PORT, nil)
}

func router(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		greet(w, r)
	case "/about":
		about(w, r)
	case "/contact":
		contact(w, r)
	default:
		notFound(w, r)
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, ini halaman utama!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ini halaman About. Kita belajar Go Webserver!")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ini halaman Contact. Hubungi kami di email@example.com")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 - Halaman tidak ditemukan")
}
