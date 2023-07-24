package httpServer

import (
	"net/http"
)

func InitHttpServer() {
	SearchPerson()
	DeletePerson()
	http.ListenAndServe(":8080", nil)
}
