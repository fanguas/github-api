package module

import (
	"os"
)

func ValidateTokenGithub() string {
	token := os.Getenv("GH_TOKEN")

	if token == "" {
		return "No se encontró el token de GitHub. Por favor, establece la variable de entorno GH_TOKEN."
	}

	return token
}
