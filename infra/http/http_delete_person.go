package httpServer

import (
	"fmt"
	"net/http"
)

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()
	var err error

	Service.DeletePerson("")

	if err != nil {
		fmt.Fprintf(w, "Erro delete: %v", err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
