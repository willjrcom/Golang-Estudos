package httpServer

import (
	"fmt"
	"net/http"
	personEntity "projetoGo/entity/person"
	"strings"
)

func SearchAllPerson(w http.ResponseWriter, r *http.Request) {
	people, _ := Service.FindAll()

	if len(people) == 0 {
		people = append(people, &personEntity.Person{})
	}

	err := modelos.ExecuteTemplate(w, "index.html", people)

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
	fmt.Println(params)
	person := personEntity.Person{}
	var err error

	_, err = Service.FindBy()

	if err != nil && !strings.Contains(err.Error(), "not found") {
		fmt.Fprintf(w, "Erro: %v", err)
	}

	err = modelos.ExecuteTemplate(w, "index.html", []personEntity.Person{person})

	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
	}
}
