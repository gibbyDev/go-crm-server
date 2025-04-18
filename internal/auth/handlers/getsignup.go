package handlers

import (
	"go-crm-server/internal/views"
	"net/http"
)

type GetSignUpHandler struct{}

func NewGetSignUpHandler() *GetSignUpHandler {
	return &GetSignUpHandler{}
}

func (h *GetSignUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := views.SignUp("Sign Up")
	err := views.Layout(c, "My website").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}