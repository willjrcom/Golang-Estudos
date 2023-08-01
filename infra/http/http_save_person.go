package httpServer

import (
	"net/http"
)

func InsertPerson(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	err := Service.RegisterPerson(r.Form)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
