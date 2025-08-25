package main

import (
	"log"
	"net/http"

	"github.com/fanguas/github-api/handlers"
)

func main() {
	log.Println("Iniciando servidor💻")
	http.HandleFunc("/api/org", handlers.OrganizacionHandler)
	http.HandleFunc("/api/org/usuario", handlers.UsuarioHandler)
	http.HandleFunc("/api/org/miembros", handlers.MiembrosHandler)
	http.HandleFunc("/api/org/repositorio", handlers.RepositorioHandler)
	http.HandleFunc("/api/org/repositorios", handlers.RepositoriosHandler)
	http.HandleFunc("/api/org/repositorios/crear", handlers.CrearRepositorioHandler)
	http.HandleFunc("/api/org/repositorios/acceso", handlers.OtorgaPermisosHandler)
	log.Println("Servidor escuchando en el puerto :8080🚀")
	log.Fatal(http.ListenAndServe(":8080", handlers.HabilitaCORS(http.DefaultServeMux)))
}
