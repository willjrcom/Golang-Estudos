package httpServer

import (
	"fmt"
	"net/http"
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
	params := r.URL.Query()
	people, err := Service.FindBy(params)

	if err != nil && !strings.Contains(err.Error(), "not found") {
		fmt.Fprintf(w, "Erro: %v", err)
	}
	fmt.Fprintf(w, "Resultado:")
	for _, person := range people {
		fmt.Fprintf(w, "%v\n", person.Name)
	}
	//err = modelos.ExecuteTemplate(w, "index.html", []personEntity.Person{})

	if err != nil {
		http.Error(w, "\nErro ao carregar template", http.StatusInternalServerError)
	}
}
