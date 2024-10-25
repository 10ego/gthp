package handlers

import (
	"net/http"

	"github.com/10ego/gthp/internal/templ/templates"
)

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := templates.Login().Render(r.Context(), w)
		if err != nil {
			h.log.Errorw("Failed to render login template", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Create a context with a timeout
	// ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	// defer cancel()

	h.log.Info("Attempting to authenticate with LDAP..")
	authenticated, err := h.ldapClient.Authenticate(r.Context(), username, password)
	if err != nil {
		h.log.Errorw("Authentication error", "error", err, "username", username)
		http.Error(w, "Authentication error", http.StatusInternalServerError)
		return
	}
	if authenticated {
		h.log.Infow("User authenticated successfully", "username", username)
		w.Write([]byte("<div>Login successful! Redirecting...</div>"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		h.log.Warnw("Invalid credentials", "username", username)
		w.Write([]byte("<div>Login failed!div>"))
		w.WriteHeader(http.StatusUnauthorized)
		templates.Login().Render(r.Context(), w)
	}
}

func (h *Handler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Implement logout logic here
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
