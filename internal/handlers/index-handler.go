package handlers

import (
	"net/http"

	"github.com/10ego/gthp/internal/templ/templates"
)

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	templates.Index(h.config.Title).Render(r.Context(), w)
}
