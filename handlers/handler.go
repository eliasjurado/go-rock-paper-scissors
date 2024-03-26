package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Message string
	}{
		"Pagina de inicio",
		"Bienvenido a Piedra, Papel o Tijera!",
	}
	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Crear Juego Nuevo")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Juego")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Jugar")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Acerca de")
}
