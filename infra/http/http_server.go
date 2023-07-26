package httpServer

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func InitHttpServer() {
	http.HandleFunc("/person/all", SearchAllPerson)
	http.HandleFunc("/person", SearchPerson)
	http.HandleFunc("/deletePerson", DeletePerson)
	http.HandleFunc("/", templateHtml)
	http.ListenAndServe(":8080", nil)
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
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(body))
	}
}
