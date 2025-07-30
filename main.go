package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/edigun-man/go-crud-api.git/handlers"
	"github.com/edigun-man/go-crud-api.git/models"
)

func main() {

	// create dummy user
	for i := range 10 {
		name := fmt.Sprintf("user%d", i)
		email := fmt.Sprintf("edi%d@example.com", i)
		models.CreateUser(name, email)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "welcome folks 🐒")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if r.URL.Query().Get("id") != "" {
				handlers.GetUser(w, r)
			} else {
				handlers.GetUsers(w, r)
			}

		case http.MethodPost:
			handlers.CreateUser(w, r)

		case http.MethodPut:
			handlers.UpdateUser(w, r)

		case http.MethodDelete:
			handlers.DeleteUser(w, r)
		}
	})

	log.Println("Server run at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
