package httpServer

import (
	"html/template"
	"net/http"
)

var modelos = template.Must(template.ParseFiles("infra/html/index.html"))

func templateHtml(w http.ResponseWriter, r *http.Request) {
	people, err := Service.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = modelos.ExecuteTemplate(w, "index.html", people)

	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
	}
}
