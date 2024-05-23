package main

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"main/core"
	"main/db"
	"main/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Path[1:]
	if name == "" {
		fmt.Fprintf(w, "Hello, World!")
	} else {
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}

func UserSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	rawBody := r.Body

	if rawBody == nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	var data models.SignupRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// timing safe comparison
	if subtle.ConstantTimeCompare([]byte(data.Password), []byte(data.ConfirmPassword)) == 0 {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	dbClient, err := core.GetDbClient()

	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, "System error", http.StatusInternalServerError)
		return
	}

	// create user
	createdUser, err := dbClient.User.CreateOne(db.User.Email.Set(data.Username)).Exec(ctx)

	if err != nil || createdUser == nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	storedPasswordHash, err := dbClient.Passwords.CreateOne(db.Passwords.User.Link(db.User.ID.Equals(createdUser.ID)), db.Passwords.PasswordHash.Set(string(hashedPassword))).Exec(ctx)

	if err != nil || storedPasswordHash == nil {
		http.Error(w, "Failed to store password", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdUser)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/signup", UserSignup)
	http.ListenAndServe(fmt.Sprintf(":%s", core.DEFAULT_PORT), nil)
}
