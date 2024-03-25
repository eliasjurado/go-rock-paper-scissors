package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/",Index)

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s", port)
	http.ListenAndServe(port, router)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hola mundo")
}