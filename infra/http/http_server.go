package httpServer

import (
	"fmt"
	"io"
	"net/http"
	"projetoGo/infra/database"
	personRepositoryGorm "projetoGo/infra/repository/gorm/person-repository"
	personService "projetoGo/usecases/person"
	"time"

	"github.com/go-chi/chi/v5"
)

var Service = personService.Service{
	Repository: &personRepositoryGorm.PersonRepository{
		Db: database.NewDb(),
	},
}

func InitHttpServer() {
	r := chi.NewRouter()
	r.Get("/person/all", SearchAllPerson)
	r.Get("/person", SearchPerson)
	r.Get("/deletePerson", DeletePerson)
	r.Post("/newPerson", InsertPerson)
	r.Get("/", templateHtml)
	http.ListenAndServe(":8080", r)
}

func getServer() {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	response, err := client.Get("https://www.google.com.br")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(body))
	}
}
