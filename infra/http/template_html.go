package httpServer

import (
	"html/template"
	"net/http"
)

var modelos = template.Must(template.ParseFiles("infra/html/index.html"))

func templateHtml(w http.ResponseWriter, r *http.Request) {
	people, _ := Service.FindAll()
	err := modelos.ExecuteTemplate(w, "index.html", people)

	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
	}
}
