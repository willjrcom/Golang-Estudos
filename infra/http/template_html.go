package httpServer

import (
	"html/template"
	"net/http"
	personRepository "projetoGo/infra/repository/person-repository"
)

var modelos = template.Must(template.ParseFiles("infra/html/index.html"))

func templateHtml(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"DeleteByName": personRepository.DeleteByName,
	}

	modelos.Funcs(funcMap)
	err := modelos.ExecuteTemplate(w, "index.html", personRepository.FindAll())

	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
	}
}
