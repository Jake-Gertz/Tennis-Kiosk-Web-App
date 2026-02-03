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
	db, err := sql.Open("sqlite", "./users.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Strings will have a string that represetns first their name and second their tension after a | as a seperator

	// Create the user info table if it does not exist
	userInfoTable := "CREATE TABLE IF NOT EXISTS userInfoTable (userID INTEGER PRIMARY KEY, stringerID INTEGER, defaultRacket TEXT, defaultMainString TEXT, defaultCrossString TEXT);"
	_, userInfoErr := db.Exec(userInfoTable)

	if userInfoErr != nil {
		log.Fatal(userInfoErr)
	}

	// Create the Rackets to pick up table
	racketsToPickUpTable := "CREATE TABLE IF NOT EXISTS racketsToPickUp (pickupID INTEGER PRIMARY KEY AUTOINCREMENT, userID INTEGER, racket TEXT, mainString TEXT, crossString TEXT, dateStrung TEXT);"
	_, racketsToPickUpErr := db.Exec(racketsToPickUpTable)

	if racketsToPickUpErr != nil {
		log.Fatal(racketsToPickUpErr)
	}

	// Create the rackets to string table
	racktsToStringTable := "CREATE TABLE IF NOT EXISTS racketsToStringTable (dropoffID INTEGER PRIMARY KEY AUTOINCREMENT, userID INTEGER, racket TEXT, mainString TEXT, crossString TEXT, dateDropedOff TEXT, stringerID INTEGER);"
	_, racketsToStringErr := db.Exec(racktsToStringTable)

	if racketsToStringErr != nil {
		log.Fatal(racketsToStringErr)
	}

	addDefaultIDs(db)

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

/* need to refactor this whole thing since we changed how the db works */
func addDefaultIDs(db *sql.DB) {
	var count int

	db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)

	if count < 3 {
		defaultAdminID := "9999"
		defaultStringerID := "8888"
		defaultUserID := "7777"

		_, err1 := db.Exec("INSERT INTO users (id) VALUES (?)", defaultAdminID)
		_, err2 := db.Exec("INSERT INTO users (id) VALUES (?)", defaultStringerID)
		_, err3 := db.Exec("INSERT INTO users (id) VALUES (?)", defaultUserID)

		if err1 != nil && err2 != nil && err3 != nil {
			log.Println("Error inserting admin ID: ", err1, err2, err3)
		}
	}
}
