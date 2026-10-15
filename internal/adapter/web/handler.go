package web

import "html/template"

type Handler struct {
	templates *template.Template
	integrationsTemplate *template.Template
}

func NewHandler(
	templates *template.Template, integrationsTemplate *template.Template,
	) *Handler {
	return &Handler{
		templates: templates, integrationsTemplate: integrationsTemplate,
	}
}
