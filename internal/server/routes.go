package server

import (
	"net/http"
	"strings"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Route for HelloWorld
	mux.HandleFunc("/", s.HelloWorldHandler)

	// Health check route
	mux.HandleFunc("/health", s.healthHandler)

	// Routes for user management
	mux.HandleFunc("/users", s.createUserHandler)        // POST: Create a user
	mux.HandleFunc("/users/all", s.getAllUsersHandler)   // GET: Retrieve all users
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if id := strings.TrimPrefix(r.URL.Path, "/users/"); id != "" {
			s.getUserByIDHandler(w, r, id) // Pass the ID to the handler
		} else {
			http.NotFound(w, r)
		}
	})

	// Routes for ticket management
	mux.HandleFunc("/tickets", s.createTicketHandler)        // POST: Create a ticket
	mux.HandleFunc("/tickets/all", s.getAllTicketsHandler)   // GET: Retrieve all tickets
	mux.HandleFunc("/tickets/", func(w http.ResponseWriter, r *http.Request) {
		if id := strings.TrimPrefix(r.URL.Path, "/tickets/"); id != "" {
			s.getTicketByIDHandler(w, r, id) // Pass the ID to the handler
		} else {
			http.NotFound(w, r)
		}
	})
	mux.HandleFunc("/tickets/update/", func(w http.ResponseWriter, r *http.Request) {
		if id := strings.TrimPrefix(r.URL.Path, "/tickets/update/"); id != "" {
			s.updateTicketHandler(w, r, id) // Pass the ID to the handler
		} else {
			http.NotFound(w, r)
		}
	})
	mux.HandleFunc("/tickets/delete/", func(w http.ResponseWriter, r *http.Request) {
		if id := strings.TrimPrefix(r.URL.Path, "/tickets/delete/"); id != "" {
			s.deleteTicketHandler(w, r, id) // Pass the ID to the handler
		} else {
			http.NotFound(w, r)
		}
	})

	return mux
}

