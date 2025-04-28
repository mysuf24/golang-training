package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var PORT = ":8080"

// Struktur data User
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/user", userHandler)

	fmt.Println("Server berjalan di http://localhost" + PORT)
	http.ListenAndServe(PORT, nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		//GET request
		user := User{
			ID:   1,
			Name: "Mocha",
		}
		json.NewEncoder(w).Encode(user)

	} else if r.Method == http.MethodPost {
		//POST request
		var newUser User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// disini cuma balikin data
		newUser.ID = 2

		json.NewEncoder(w).Encode(newUser)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
