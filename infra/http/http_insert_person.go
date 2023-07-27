package httpServer

import (
	"net/http"
	personEntity "projetoGo/entity/person"
	personRepository "projetoGo/infra/repository/person-repository"
	"time"
)

func InsertPerson(w http.ResponseWriter, r *http.Request) {
	layout := "2023-01-01 00:00:00"

	r.ParseForm()
	name := r.FormValue("name")
	cpf := r.FormValue("cpf")
	genre := r.FormValue("genre")
	date, _ := time.Parse(layout, r.FormValue("date"))

	person, _ := personEntity.NewPersonBuilder(name, cpf).WithBirthday(date).WithGenre(genre).Build()

	personRepository.InsertPerson(person)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
