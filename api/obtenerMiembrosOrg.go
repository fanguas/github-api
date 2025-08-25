package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fanguas/github-api/module"
)

func ObtenerMiembrosOrganizacion(org string) ([]Miembro, error) {
	if org == "" {
		return nil, fmt.Errorf("el nombre de la organización es obligatorio")
	}

	token := module.ValidateTokenGithub()
	var todosLosMiembros []Miembro
	cliente := &http.Client{}
	pagina := 1
	elementosPorPag := 30

	for {
		urlPaginada := fmt.Sprintf(GitHubAPIBaseURL+"/orgs/%s/members?page=%d&per_page=%d", org, pagina, elementosPorPag)

		solicitud, err := http.NewRequest(http.MethodGet, urlPaginada, nil)
		if err != nil {
			return nil, fmt.Errorf("error al crear la solicitud: %v", err)
		}

		solicitud.Header.Set("Accept", HeaderAccept)
		solicitud.Header.Set(HeaderAuth, "Bearer "+token)

		respuesta, err := cliente.Do(solicitud)
		if err != nil {
			return nil, fmt.Errorf("error al realizar la solicitud: %v", err)
		}
		defer respuesta.Body.Close()

		if respuesta.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error al obtener los miembros: código %d", respuesta.StatusCode)
		}

		cuerpo, err := io.ReadAll(respuesta.Body)
		if err != nil {
			return nil, fmt.Errorf("error al leer la respuesta: %v", err)
		}

		var miembros []Miembro
		err = json.Unmarshal(cuerpo, &miembros)
		if err != nil {
			return nil, fmt.Errorf("error al deserializar la respuesta: %v", err)
		}

		todosLosMiembros = append(todosLosMiembros, miembros...)

		if len(miembros) < elementosPorPag {
			break
		}
		pagina++
	}

	return todosLosMiembros, nil
}
