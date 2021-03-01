package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/users", handleUsers)
	http.ListenAndServe(":8080", nil)
}

var users = []User{
	{1, "a"},
	{2, "b"},
	{3, "c"},
}

type User struct {
	ID   int
	Name string
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		usersJson, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(usersJson)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w,"{\"message\": \"not found\"}")

	}
}

