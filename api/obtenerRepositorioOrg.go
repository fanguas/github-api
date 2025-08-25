package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fanguas/github-api/module"
)

func ObtenerRepositorio(org, repo string) (*Repositorio, []Miembro, error) {
	if org == "" || repo == "" {
		return nil, nil, fmt.Errorf("organización y repositorio son obligatorios")
	}

	token := module.ValidateTokenGithub()
	cliente := &http.Client{}
	urlRepo := fmt.Sprintf(GitHubAPIBaseURL+"/repos/%s/%s", org, repo)
	urlColaboradores := fmt.Sprintf(GitHubAPIBaseURL+"/repos/%s/%s/collaborators", org, repo)

	// Función auxiliar para hacer GET y decodificar JSON
	fetchURL := func(url string, target interface{}) error {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return fmt.Errorf("error creando solicitud: %w", err)
		}
		req.Header.Set("Accept", HeaderAccept)
		req.Header.Set(HeaderAuth, "Bearer "+token)

		resp, err := cliente.Do(req)
		if err != nil {
			return fmt.Errorf("error haciendo solicitud HTTP: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("error en la respuesta: %s", resp.Status)
		}

		if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
			return fmt.Errorf("error decodificando JSON: %w", err)
		}
		return nil
	}

	var repositorio Repositorio
	if err := fetchURL(urlRepo, &repositorio); err != nil {
		return nil, nil, fmt.Errorf("error obteniendo repositorio: %w", err)
	}

	var colaboradores []Miembro
	if err := fetchURL(urlColaboradores, &colaboradores); err != nil {
		return &repositorio, nil, fmt.Errorf("error obteniendo colaboradores: %w", err)
	}

	return &repositorio, colaboradores, nil
}
