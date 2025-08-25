package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fanguas/github-api/module"
)

func GenerarRepositorio(org, nombre, descripcion string, privado bool) error {
	if org == "" || nombre == "" || descripcion == "" {
		return fmt.Errorf("los parámetros 'org', 'nombre' y 'descripcion' son obligatorios")
	}

	token := module.ValidateTokenGithub()
	cliente := &http.Client{}
	url := fmt.Sprintf(GitHubAPIBaseURL+"/orgs/%s/repos", org)

	body := map[string]interface{}{
		"name":        nombre,
		"description": descripcion,
		"private":     privado,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error al serializar el cuerpo de la solicitud: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return fmt.Errorf("error al crear la solicitud: %v", err)
	}

	req.Header.Set("Accept", HeaderAccept)
	req.Header.Set(HeaderAuth, "Bearer "+token)
	req.Header.Set(HeaderContentType, "application/json")

	res, err := cliente.Do(req)
	if err != nil {
		return fmt.Errorf("error al enviar la solicitud: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("error al crear el repositorio: código %d", res.StatusCode)
	}

	fmt.Printf("Repositorio '%s' creado exitosamente en la organización '%s'\n", nombre, org)

	return nil
}
