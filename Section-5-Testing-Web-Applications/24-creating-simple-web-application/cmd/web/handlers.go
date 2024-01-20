package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

type TemplateData struct {
	IP   string
	Data map[string]any
}

var pathToTemplates = "./templates/"

func (a *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = a.render(w, r, "home.page.gohtml", &TemplateData{})
}

func (a *application) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	log.Println(email, password)
	_, err = fmt.Fprintf(w, email)
	if err != nil {
		return
	}
}

func (a *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	// parse template from disk
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	data.IP = a.ipFromContext(r.Context())

	//execute template, passing data if any
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}