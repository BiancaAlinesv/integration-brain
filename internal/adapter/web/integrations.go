package web

import "net/http"

func (h *Handler) Integrations(w http.ResponseWriter, r *http.Request) {
	err := h.integrationsTemplate.ExecuteTemplate(w, "integrations-content", nil)
	if err != nil {
		http.Error(w, "Erro ao renderizar página", http.StatusInternalServerError)
		return
	}
}