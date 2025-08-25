package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fanguas/github-api/module"
)

func OtorgaAccesoAMiembro(org string, repositorios []string, miembro, permiso string) error {
	if org == "" || miembro == "" || permiso == "" {
		return fmt.Errorf("el token, el nombre de la organización, el miembro y el permiso son obligatorios")
	}

	token := module.ValidateTokenGithub()
	cliente := &http.Client{}

	for _, repo := range repositorios {
		url := fmt.Sprintf(GitHubAPIBaseURL+"/repos/%s/%s/collaborators/%s", org, repo, miembro)

		body := map[string]string{"permission": permiso}
		bodyJSON, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("error al serializar el permiso: %v", err)
		}

		req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bodyJSON))
		if err != nil {
			return fmt.Errorf("error al crear la solicitud para repo %s: %v", repo, err)
		}

		req.Header.Set("Accept", HeaderAccept)
		req.Header.Set(HeaderAuth, "Bearer "+token)
		req.Header.Set(HeaderContentType, "application/json")

		resp, err := cliente.Do(req)
		if err != nil {
			return fmt.Errorf("error al enviar solicitud para repo %s: %v", repo, err)
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error al leer respuesta para repo %s: %v", repo, err)
		}

		if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
			return fmt.Errorf("error al otorgar acceso a %s en repo %s: código %d, respuesta: %s", miembro, repo, resp.StatusCode, string(respBody))
		}
	}

	fmt.Printf("Permiso '%s' otorgado a %s en los repositorios: %v\n", permiso, miembro, repositorios)
	return nil
}
