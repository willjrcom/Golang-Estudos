package httpServer

import (
	"fmt"
	"net/http"
	personRepository "projetoGo/infra/repository/person-repository"
)

func DeletePerson() {
	http.HandleFunc("/deletePerson", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		var err error

		if name := params.Get("name"); len(name) != 0 {
			err = personRepository.DeleteByName(name)
		}

		if err != nil {
			fmt.Fprintf(w, "Erro delete: %v", err)
		}

		fmt.Fprintf(w, "Delete 200")
	})
}