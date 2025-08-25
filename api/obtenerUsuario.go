package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fanguas/github-api/module"
)

func ObtenerUsuario(usuario string) (*Usuario, error) {
	if usuario == "" {
		return nil, fmt.Errorf("el usuario es obligatorio")
	}

	token := module.ValidateTokenGithub()
	url := fmt.Sprintf(GitHubAPIBaseURL+"/users/%s", usuario)
	var usuarioData Usuario
	cliente := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creando solicitud: %w", err)
	}

	req.Header.Set("Accept", HeaderAccept)
	req.Header.Set(HeaderAuth, "Bearer "+token)

	respuesta, err := cliente.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error haciendo solicitud HTTP: %w", err)
	}
	defer respuesta.Body.Close()

	if respuesta.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la respuesta: %s", respuesta.Status)
	}
	err = json.NewDecoder(respuesta.Body).Decode(&usuarioData)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta: %w", err)
	}
	return &usuarioData, nil
}
