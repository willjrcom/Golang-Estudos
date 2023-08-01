package httpServer

import (
	"net/http"
)

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()

	r.ParseForm()
	err := Service.DeletePerson(r.Form)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
