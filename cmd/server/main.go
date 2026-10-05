package main

import (
	"html/template"
	"integration-brain/internal/adapter/web"
	"log"
	"net/http" 
)

func main() {
	templates, err := template.ParseFiles(
		"web/templates/layouts/base.html",
		"web/templates/partials/header.html",
		"web/templates/partials/navigation.html",
		"web/templates/partials/footer.html",
		"web/templates/pages/home.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	handler := web.NewHandler(templates)

	log.Println("Iniciando servidor na porta 8080...")
	fileServer := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", handler.Home)
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}


