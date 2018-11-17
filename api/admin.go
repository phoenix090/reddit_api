package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

/************** Admin handlers with username and pwd authentication ******************/

// GetUser retrieves one user from the db.
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	pwd := vars["pwd"]
	if username != os.Getenv("ADMIN_USR") || pwd != os.Getenv("ADMIN_PWD") {
		http.Error(w, "ADMIN ACCESS ONLY!!", http.StatusBadRequest)
		return
	}
	id := vars["id"]
	user, err := globalDB.Get(id)
	if err != nil {
		http.Error(w, "Could not get user", http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}

// GetAllUsers gets all user registered in the db.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	username := vars["username"]
	pwd := vars["pwd"]
	if username != os.Getenv("ADMIN_USR") || pwd != os.Getenv("ADMIN_PWD") {
		http.Error(w, "ADMIN ACCESS ONLY!!", http.StatusBadRequest)
		return
	}

	users, err := globalDB.GetAll()
	if err != nil {
		http.Error(w, "Could not get users", http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

}
