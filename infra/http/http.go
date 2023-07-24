package httpServer

import (
	"fmt"
	"net/http"
	person_repository "projetoGo/infra/repository/person-repository"
)

func initHttpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.HandleFunc("/person", func(w http.ResponseWriter, r *http.Request) {
		person := person_repository.FindAll()

		fmt.Fprintf(w, "%v", person)
	})
	http.ListenAndServe(":8080", nil)
}
