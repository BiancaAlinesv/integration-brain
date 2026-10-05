package web

import (
	"net/http"
	"log"
)
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	err := h.templates.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		log.Printf("Erro ao renderizar página inicial: %v", err)
		http.Error(w, "Erro ao renderizar página", http.StatusInternalServerError)
		return
	}
}