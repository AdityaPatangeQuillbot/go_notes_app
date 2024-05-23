package handlers

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

func HelloHandler(w http.ResponseWriter, r *http.Request) {
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

	dbClient, err := core.GetDbClient()

	if err != nil {
		core.FormatErrorResponseJSON(&w, fmt.Errorf("database connection error"), http.StatusInternalServerError)
		return
	}

	if rawBody == nil {
		core.FormatErrorResponseJSON(&w, fmt.Errorf("invalid request body"), http.StatusBadRequest)
		return
	}

	var data models.SignupRequest

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		core.FormatErrorResponseJSON(&w, fmt.Errorf("invalid request body"), http.StatusBadRequest)
		return
	}

	foundUser, err := dbClient.User.FindFirst(db.User.Email.Equals(data.Username)).Exec(ctx)

	if err != nil {
		core.FormatErrorResponseJSON(&w, fmt.Errorf("system error"), http.StatusInternalServerError)
		return
	}

	if foundUser != nil {
		core.FormatErrorResponseJSON(&w, fmt.Errorf("user already exists"), http.StatusConflict)
		return
	}

	_, err = core.ValidateSignupRequest(data)

	if err != nil {
		core.FormatErrorResponseJSON(&w, err, http.StatusBadRequest)
		return
	}

	// timing safe comparison
	if subtle.ConstantTimeCompare([]byte(data.Password), []byte(data.ConfirmPassword)) == 0 {
		core.FormatErrorResponseJSON(&w, fmt.Errorf("passwords do not match"), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		core.FormatErrorResponseJSON(&w, fmt.Errorf("system error"), http.StatusInternalServerError)
		return
	}

	// create user
	createdUser, err := dbClient.User.CreateOne(db.User.Email.Set(data.Username)).Exec(ctx)

	if err != nil || createdUser == nil {
		core.FormatErrorResponseJSON(&w, fmt.Errorf("failed to create user"), http.StatusInternalServerError)
		return
	}

	storedPasswordHash, err := dbClient.Passwords.CreateOne(db.Passwords.User.Link(db.User.ID.Equals(createdUser.ID)), db.Passwords.PasswordHash.Set(string(hashedPassword))).Exec(ctx)

	if err != nil || storedPasswordHash == nil {
		core.FormatErrorResponseJSON(&w, fmt.Errorf("failed to store password"), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdUser)
}
