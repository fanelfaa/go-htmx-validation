package controllers

import (
	"fmt"
	"go-htmx-form-validation/form"
	"go-htmx-form-validation/templates"
	"html/template"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UsersControllers struct{}

type FormUser struct {
	Name    string `form:"name" validate:"required"`
	Email   string `form:"email" validate:"required,email"`
	Age     int8   `form:"age" validate:"required,gte=5,lte=80"`
	Address string `form:"address"`
	Active  bool   `form:"active"`
}

func (uc UsersControllers) Add(w http.ResponseWriter, r *http.Request) {
	template.Must(
		template.ParseFS(
			templates.FS,
			"users/add.html",
			"users/form.html",
		),
	).Execute(w, nil)
}

type FormErrors map[string]string

func (uc UsersControllers) Store(w http.ResponseWriter, r *http.Request) {
	// check if request is not from hx then error
	if r.Header.Get("HX-Request") != "true" {
		http.Error(w, "Only HTMX Request", http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	values := r.Form

	var formUser FormUser

	err = form.Decoder.Decode(&formUser, values)
	if err != nil {
		fmt.Println(err)
	}

	err = form.Validate.Struct(formUser)

	if err != nil {
		errors := make(FormErrors)

		for _, err := range err.(validator.ValidationErrors) {
			msg, errMap := form.MapValidationError(err)
			if errMap == nil {
				errors[err.StructField()] = msg
			} else {
				errors[err.StructField()] = err.Tag()
			}
		}

		template.Must(
			template.ParseFS(
				templates.FS,
				"users/form.html",
			),
		).ExecuteTemplate(w, "form-fields", map[string]interface{}{
			"Values": formUser,
			"Errors": errors,
		})
		return
	}
	w.Header().Add("HX-Redirect", "/")
}
