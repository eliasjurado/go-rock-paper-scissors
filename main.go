package main

import (
	"fmt"
	"net/http"
	"rpsweb/handlers"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s", port)
	http.ListenAndServe(port, router)
}
