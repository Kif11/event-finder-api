package junk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type db struct {
	users map[int]user
}

func (d *db) populateUsers() {
	users := make(map[int]user)

	users[0] = user{"test@mail.com", "12345"}
	users[1] = user{"kif@mail.com", "8888888"}

	d.users = users
}

func (d *db) getUser(id int) (user, error) {
	if val, ok := d.users[id]; ok {
		return val, nil
	}

	return user{}, fmt.Errorf("db: user with %d id doesn't exists", id)
}

var newDB db

func respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var buf bytes.Buffer

	// Encode data to json first
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write data that we want to send
	if _, err := io.Copy(w, &buf); err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val, ok := r.Header["Authorization"]

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokens := strings.Split(val[0], " ")

		if len(tokens) < 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if tokens[1] != "JC7UGUY35V6ICPGOQ7MJ" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var u user
	err := decoder.Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(u.Email + " " + u.Password)

	w.WriteHeader(http.StatusOK)
}

var getUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		respond(w, r, http.StatusInternalServerError, err)
		return
	}

	user, err := newDB.getUser(userID)
	if err != nil {
		respond(w, r, http.StatusInternalServerError, err.Error())
		// w.WriteHeader(500)
		return
	}

	respond(w, r, http.StatusOK, user)
})

func main() {
	r := mux.NewRouter()
	newDB = db{}
	newDB.populateUsers()

	r.HandleFunc("/user", createUser).Methods("POST")
	r.Handle("/user/{id}", authMiddleware(getUser)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
