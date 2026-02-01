/* lots to modify here when making the table you actually want*/
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

type UserRequest struct {
	UserID string `json:"userId"`
}

type UserResponse struct {
	Exists bool `json:"exists"`
}

func main() {
	// 1. Open (or create) the SQLite database
	db, err := sql.Open("sqlite", "./users.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. Create the table if it doesn't exist
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users (id TEXT PRIMARY KEY)")
	statement.Exec()

	addAdminIDs(db)

	// 3. Define the API endpoint
	http.HandleFunc("/check-user", func(w http.ResponseWriter, r *http.Request) {
		// Allow the Vue app to talk to this server (CORS)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		var req UserRequest
		json.NewDecoder(r.Body).Decode(&req)

		var id string
		err := db.QueryRow("SELECT id FROM users WHERE id = ?", req.UserID).Scan(&id)

		exists := err == nil
		json.NewEncoder(w).Encode(UserResponse{Exists: exists})
	})

	fmt.Println("Go server starting at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func addAdminIDs(db *sql.DB) {
	var count int

	db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)

	if count == 0 {
		defaultAdminID := "9999"

		_, err := db.Exec("INSERT INTO users (id) VALUES (?)", defaultAdminID)

		if err != nil {
			log.Println("Error inserting admin ID: ", err)
		}
	}
}
