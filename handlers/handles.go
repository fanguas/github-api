package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fanguas/github-api/api"
)

func OrganizacionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	org := r.URL.Query().Get("org")
	if org == "" {
		http.Error(w, "El parámetro 'org' es requerido", http.StatusBadRequest)
		return
	}

	organizacion, variables, secretos, err := api.ObtenerOrganizacion(org)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"organizacion": organizacion,
		"variables":    variables,
		"secretos":     secretos,
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	user := r.URL.Query().Get("usuario")
	if user == "" {
		http.Error(w, "El parámetro 'usuario' es requerido", http.StatusBadRequest)
		return
	}

	usuario, err := api.ObtenerUsuario(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}

func MiembrosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	org := r.URL.Query().Get("org")
	if org == "" {
		http.Error(w, "El parámetro 'org' es requerido", http.StatusBadRequest)
		return
	}

	miembros, err := api.ObtenerMiembrosOrganizacion(org)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(miembros)
}

func RepositorioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	org := r.URL.Query().Get("org")
	repo := r.URL.Query().Get("repo")
	if org == "" || repo == "" {
		http.Error(w, "Los parámetros 'org' y 'repo' son requeridos", http.StatusBadRequest)
		return
	}

	repositorio, colaboradores, err := api.ObtenerRepositorio(org, repo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := api.RespuestaRepositorio{
		Repositorio:   repositorio,
		Colaboradores: colaboradores,
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func RepositoriosHandler(w http.ResponseWriter, r *http.Request) {
	org := r.URL.Query().Get("org")
	if org == "" {
		http.Error(w, "El parámetro 'org' es requerido", http.StatusBadRequest)
		return
	}

	repos, err := api.ObtenerRepositoriosOrganizacion(org)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repos)
}

func CrearRepositorioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req api.RepositorioRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Error al parsear JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if req.Org == "" || req.Nombre == "" || req.Descripcion == "" {
		http.Error(w, "Los campos 'org', 'nombre' y 'descripcion' son obligatorios", http.StatusBadRequest)
		return
	}

	if err := api.GenerarRepositorio(req.Org, req.Nombre, req.Descripcion, req.Privado); err != nil {
		http.Error(w, "Error interno al crear repositorio: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Repositorio creado exitosamente 📦",
	})
}

func OtorgaPermisosHandler(w http.ResponseWriter, r *http.Request) {
	org := r.URL.Query().Get("org")
	if org == "" {
		http.Error(w, "El parámetro 'org' es requerido", http.StatusBadRequest)
		return
	}

	usuario := r.URL.Query().Get("miembro")
	if usuario == "" {
		http.Error(w, "El parámetro 'miembro' es requerido", http.StatusBadRequest)
		return
	}

	repos := r.URL.Query()["repositorio"]
	if len(repos) == 0 {
		http.Error(w, "Debes enviar al menos un 'repositorio'", http.StatusBadRequest)
		return
	}

	permiso := r.URL.Query().Get("permiso")
	if permiso == "" {
		http.Error(w, "El parámetro 'permiso' es requerido", http.StatusBadRequest)
		return
	}

	err := api.OtorgaAccesoAMiembro(org, repos, usuario, permiso)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func HabilitaCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
