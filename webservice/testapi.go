package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PORT = ":8080"
	db   *gorm.DB // variabel global database
)

// Struktur data User
type User struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func main() {
	var err error

	// Koneksi ke PostgreSQL
	dsn := "host=localhost user=postgres password=123 dbname=mydb port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek ke database:", err)
	}

	// Auto migrate table
	db.AutoMigrate(&User{})

	// Routing
	http.HandleFunc("/user", userHandler)

	fmt.Println("Server berjalan di http://localhost" + PORT)
	http.ListenAndServe(PORT, nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// GET request - ambil semua user
	if r.Method == http.MethodGet {
		var users []User
		result := db.Find(&users) // Mengambil semua user
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users) // kirim semua user dalam bentuk JSON

	} else if r.Method == http.MethodPost {
		// POST request - simpan user baru
		var newUser User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Simpan ke database
		result := db.Create(&newUser)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(newUser)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
