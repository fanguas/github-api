package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fanguas/github-api/module"
)

func ObtenerOrganizacion(org string) (*Organizacion, *Variables, *Secretos, error) {
	if org == "" {
		return nil, nil, nil, fmt.Errorf("la organización es obligatoria")
	}

	token := module.ValidateTokenGithub()
	cliente := &http.Client{}
	urlOrg := fmt.Sprintf(GitHubAPIBaseURL+"/orgs/%s", org)
	urlVars := fmt.Sprintf(GitHubAPIBaseURL+"/orgs/%s/actions/variables", org)
	urlSecrets := fmt.Sprintf(GitHubAPIBaseURL+"/orgs/%s/actions/secrets", org)

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

	// Obtener datos de la organización
	var organizacion Organizacion
	if err := fetchURL(urlOrg, &organizacion); err != nil {
		return nil, nil, nil, fmt.Errorf("error obteniendo organización: %w", err)
	}

	// Obtener variables de la organización
	var variables Variables
	if err := fetchURL(urlVars, &variables); err != nil {
		return &organizacion, nil, nil, fmt.Errorf("error obteniendo variables: %w", err)
	}

	// Obtener secretos de la organización
	var secretos Secretos
	if err := fetchURL(urlSecrets, &secretos); err != nil {
		return &organizacion, &variables, nil, fmt.Errorf("error obteniendo secretos: %w", err)
	}

	return &organizacion, &variables, &secretos, nil
}
