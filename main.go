package main

import (
	"fmt"
	"go-htmx-form-validation/controllers"
	"go-htmx-form-validation/form"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type FormAddRole struct {
	Name   string
	Errors map[string]string
}

func main() {
	form.InitDecoder()
	form.InitValidate()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	usersC := controllers.UsersControllers{}

	r.Route("/users", func(r chi.Router) {
		r.Get("/add", usersC.Add)
		r.Post("/add", usersC.Store)
	})

	fmt.Println("Server start at :3000")
	http.ListenAndServe(":3000", r)
}
