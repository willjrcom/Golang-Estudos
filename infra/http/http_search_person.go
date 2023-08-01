package httpServer

import (
	"fmt"
	"net/http"
	personEntity "projetoGo/entity/person"
	"strings"
)

func SearchAllPerson(w http.ResponseWriter, r *http.Request) {
	people, err := Service.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = modelos.ExecuteTemplate(w, "index.html", people)

	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
	}
}

func SearchPerson(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()) == 0 {
		fmt.Fprintf(w, "Hello World! %v", r.URL.User.Username())
		return
	}

	params := r.URL.Query()
	_, err := Service.FindBy(params)

	if err != nil && !strings.Contains(err.Error(), "not found") {
		fmt.Fprintf(w, "Erro: %v", err)
	}

	err = modelos.ExecuteTemplate(w, "index.html", []personEntity.Person{})

	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
	}
}
