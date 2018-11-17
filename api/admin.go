package api

import (
	"encoding/json"
	"fmt"
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

	// Checking if the admin creds are correct, otherwise 400
	if username != os.Getenv("ADMIN_USR") || pwd != os.Getenv("ADMIN_PWD") {
		// StatusForbidden would make sense too her
		http.Error(w, "ADMIN ACCESS ONLY!!", http.StatusUnauthorized)
		return
	}
	id := vars["id"]
	user, err := globalDB.Get(id)
	if err != nil {
		http.Error(w, "Could not get user", http.StatusNotFound)
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
		http.Error(w, "ADMIN ACCESS ONLY!!", http.StatusUnauthorized)
		return
	}

	users, err := globalDB.GetAll()
	if err != nil {
		http.Error(w, "Could not get users", http.StatusNotFound)
		return
	}

	if err = json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}

// DeleteOneUser deletes one user from db
func DeleteOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	username := vars["username"]
	pwd := vars["pwd"]
	id := vars["id"]

	if username != os.Getenv("ADMIN_USR") || pwd != os.Getenv("ADMIN_PWD") {
		http.Error(w, "ADMIN ACCESS ONLY!!", http.StatusUnauthorized)
		return
	}

	err := globalDB.DeleteUser(id)
	if err != nil {
		http.Error(w, "Did't find user", http.StatusNotModified)
		return
	}

	fmt.Fprint(w, "OK")
}

// DeleteAllUsers deletes all the users
func DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	username := vars["username"]
	pwd := vars["pwd"]

	if username != os.Getenv("ADMIN_USR") || pwd != os.Getenv("ADMIN_PWD") {
		http.Error(w, "ADMIN ACCESS ONLY!!", http.StatusUnauthorized)
		return
	}

	err := globalDB.DeleteAll()
	if err != nil {
		http.Error(w, "Error occurred, could't delete users", http.StatusNotModified)
		return
	}

	fmt.Fprint(w, "All the users have been deleted")
}
