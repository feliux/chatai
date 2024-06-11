package handler

import (
	"net/http"

	"github.com/feliux/chatai/view/home"
)

// Render the landing page from view/home/index.templ.
func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}
