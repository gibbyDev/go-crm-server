package handlers

import (
	"go-crm-server/internal/views"
	"net/http"
)

type GetResetHandler struct{}

func NewGetResetHandler() *GetResetHandler {
	return &GetResetHandler{}
}

func (h *GetResetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := views.Reset("Password Reset")
	err := views.Layout(c, "My website").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}