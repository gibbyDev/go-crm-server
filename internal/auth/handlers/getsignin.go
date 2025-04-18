package handlers

import (
	"go-crm-server/internal/views"
	"net/http"
)

type GetSignInHandler struct{}

func NewGetSignInHandler() *GetSignInHandler {
	return &GetSignInHandler{}
}

func (h *GetSignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := views.SignIn("Sign In")
	err := views.Layout(c, "My website").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}