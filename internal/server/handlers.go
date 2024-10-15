package server

import (
 	"io/ioutil"
	"encoding/json"
	"log"
	"net/http"
)

// User struct for demo purposes, replace this with your actual database model
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// Ticket struct for demo purposes, replace this with your actual database model
type Ticket struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// HelloWorldHandler returns a Hello World message as JSON
func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"message": "Hello World",
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

// HealthHandler returns the health of the system as JSON
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

// UsersHandler handles the creation of a new user
func (s *Server) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Read the JSON request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Unmarshal JSON into a User struct
	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Normally, here you'd write to the database using s.db.CreateUser(user)
	log.Printf("User created: %+v", user)

	// Respond with the created user in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetAllUsersHandler retrieves all users (simulating with an array for now)
func (s *Server) getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Normally, you would fetch from the database
	users := []User{
		{ID: "1", Name: "John Doe", Email: "john@example.com", Role: "admin"},
		{ID: "2", Name: "Jane Smith", Email: "jane@example.com", Role: "user"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByIDHandler retrieves a user by its ID
func (s *Server) getUserByIDHandler(w http.ResponseWriter, r *http.Request, id string) {
	// Use the ID in the logic here
	// Example: Fetch user from the database using id
	log.Printf("Accessed /users/%s route", id)
	w.WriteHeader(http.StatusOK)
}

// UpdateUserHandler updates a user by their ID
func (s *Server) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Read and parse JSON request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Unmarshal the JSON into a User struct
	var updatedUser User
	if err := json.Unmarshal(body, &updatedUser); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Normally, you would update the user in the database using their ID
	log.Printf("User updated: %+v", updatedUser)

	// Respond with the updated user in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUserHandler deletes a user by their ID
func (s *Server) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL (normally you'd use a router for this)
	// For this example, let's assume we're deleting user ID "1"
	userID := "1"

	// Normally, you would delete the user from the database
	log.Printf("User deleted with ID: %s", userID)

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted", "userID": userID})
}


// CreateTicketHandler handles the creation of a new ticket
func (s *Server) createTicketHandler(w http.ResponseWriter, r *http.Request) {
	// Read the JSON request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Unmarshal JSON into a Ticket struct
	var ticket Ticket
	if err := json.Unmarshal(body, &ticket); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Normally, here you'd write to the database using s.db.CreateTicket(ticket)
	log.Printf("Ticket created: %+v", ticket)

	// Respond with the created ticket in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ticket)
}

// GetAllTicketsHandler retrieves all tickets (simulating with an array for now)
func (s *Server) getAllTicketsHandler(w http.ResponseWriter, r *http.Request) {
	// Normally, you would fetch from the database
	tickets := []Ticket{
		{ID: "1", Title: "Fix bug", Status: "open"},
		{ID: "2", Title: "Add feature", Status: "closed"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

// GetTicketByIDHandler retrieves a ticket by its ID
func (s *Server) getTicketByIDHandler(w http.ResponseWriter, r *http.Request, id string) {
	// Use the ID in the logic here
	log.Printf("Accessed /tickets/%s route", id)
	w.WriteHeader(http.StatusOK)
}

// UpdateTicketHandler updates a ticket by its ID
func (s *Server) updateTicketHandler(w http.ResponseWriter, r *http.Request, id string) {
	// Use the ID in the logic here to update the ticket
	log.Printf("Updated ticket with ID: %s", id)
	w.WriteHeader(http.StatusOK)
}

// DeleteTicketHandler deletes a ticket by its ID
func (s *Server) deleteTicketHandler(w http.ResponseWriter, r *http.Request, id string) {
	// Use the ID in the logic here to delete the ticket
	log.Printf("Deleted ticket with ID: %s", id)
	w.WriteHeader(http.StatusOK)
}
