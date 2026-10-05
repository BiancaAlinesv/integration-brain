package web

import "html/template"

type Handler struct {
	templates *template.Template
}

func NewHandler(templates *template.Template) *Handler {
	return &Handler{
		templates: templates,
	}
}
