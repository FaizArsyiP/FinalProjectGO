package main

import (
	"fmt"
	"net/http"

	"github.com/FaizArsyiP/FINALPROJECT/src/handler"
	"github.com/gorilla/mux"
)

func main() {
	// Menggunakan Gorilla Mux sebagai router
	r := mux.NewRouter()

	// Route dasar untuk mengecek server
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from server"))
	})

	// Route untuk buku
	r.HandleFunc("/api/buku", handler.BookHandler).Methods("GET", "POST")
	r.HandleFunc("/api/buku/{id}", handler.BookHandler).Methods("GET", "PUT", "DELETE")

	// Route untuk karyawan
	r.HandleFunc("/api/karyawan", handler.KaryawanHandler)

	// Jalankan server di port 8080
	fmt.Println("Server is running on port 3500")
	err := http.ListenAndServe(":3500", r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
