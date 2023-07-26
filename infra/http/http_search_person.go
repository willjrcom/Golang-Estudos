package httpServer

import (
	"fmt"
	"net/http"
	personEntity "projetoGo/entity/person"
	personRepository "projetoGo/infra/repository/person-repository"
)

func SearchAllPerson(w http.ResponseWriter, r *http.Request) {
	person := personRepository.FindAll()

	fmt.Fprintf(w, "%v - %v", person, r.URL.User.Username())
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

	if name := params.Get("name"); len(name) != 0 {
		person, err = personRepository.FindByName(name)

	} else if cpf := params.Get("cpf"); len(cpf) != 0 {
		fmt.Println(cpf)
		person, err = personRepository.FindByCpf(cpf)
	}

	if err != nil {
		fmt.Fprintf(w, "Erro: %v", err)
	}
	fmt.Fprintf(w, "%v", person)
}
