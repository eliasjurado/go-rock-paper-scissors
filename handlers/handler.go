package handlers

import (
	"net/http"
	"text/template"
)

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, nil, "index")
}

func RenderTemplate(w http.ResponseWriter, data any, page string) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page+".html"))
	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, nil, "newgame")
}

func Game(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, nil, "game")
}

func Play(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, nil, "play")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, nil, "about")
}
