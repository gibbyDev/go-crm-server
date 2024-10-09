package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HelloWorldHandler)
	mux.HandleFunc("/health", s.healthHandler)

	// routes for user management
	mux.HandleFunc("/users", s.usersHandler)          
	mux.HandleFunc("/users/all", s.allUsersHandler)   
	mux.HandleFunc("/users/", s.userByIDHandler)     

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

// UsersHandler logs when the /users route is accessed
func (s *Server) usersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Accessed /users route")
	w.WriteHeader(http.StatusOK) // Respond with a 200 OK status
}

// AllUsersHandler logs when the /users/all route is accessed
func (s *Server) allUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Accessed /users/all route")
	w.WriteHeader(http.StatusOK) // Respond with a 200 OK status
}

// UserByIDHandler logs when the /users/{id} route is accessed
func (s *Server) userByIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Accessed /users/{id} route")
	w.WriteHeader(http.StatusOK) // Respond with a 200 OK status
}
