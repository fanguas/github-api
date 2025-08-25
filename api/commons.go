package api

// Constantes para interactuar con la API de GitHub
const (
	GitHubAPIBaseURL  = "https://api.github.com"
	HeaderAccept      = "Accept"
	HeaderAuth        = "Authorization"
	HeaderContentType = "Content-Type"
)

// Organizacion representa una organización de GitHub.
type Organizacion struct {
	ID                   int    `json:"id"`
	Login                string `json:"login"`
	Nombre               string `json:"name"`
	HTMLURL              string `json:"html_url"`
	Colaboradores        int    `json:"collaborators"`
	Seguidores           int    `json:"followers"`
	AvatarURL            string `json:"avatar_url"`
	RepositoriosPublicos int    `json:"public_repos"`
	RepositoriosPrivados int    `json:"total_private_repos"`
}

// Miembro representa un miembro de una organización.
type Miembro struct {
	ID        int    `json:"id"`
	Alias     string `json:"login"`
	HTMLURL   string `json:"html_url"`
	AvatarURL string `json:"avatar_url"`
}

// Variables representa las variables de una organización.
type Variables struct {
	Variables []struct {
		Nombre string `json:"name"`
		Value  string `json:"value"`
	} `json:"variables"`
}

// Secretos representa los secretos de una organización.
type Secretos struct {
	Secretos []struct {
		Nombre        string `json:"name"`
		FechaCreacion string `json:"created_at"`
	} `json:"secrets"`
}

// Usuario representa un usuario individual de GitHub.
type Usuario struct {
	ID                   int    `json:"id"`
	Login                string `json:"login"`
	Nombre               string `json:"name"`
	Compania             string `json:"company"`
	Ubicacion            string `json:"location"`
	HTMLURL              string `json:"html_url"`
	AvatarURL            string `json:"avatar_url"`
	RepositoriosPublicos int    `json:"public_repos"`
	RepositoriosPrivados int    `json:"total_private_repos"`
}

// Repositorio representa un repositorio en GitHub.
type Repositorio struct {
	ID            int         `json:"id"`
	Nombre        string      `json:"name"`
	DBranch       string      `json:"default_branch"`
	Lenguaje      string      `json:"language"`
	HTMLURL       string      `json:"html_url"`
	Descripcion   string      `json:"description"`
	FechaCreacion string      `json:"created_at"`
	Propietario   Propietario `json:"owner"`
}

// Propietario representa el dueño de un repositorio.
type Propietario struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

// RepositorioResponse representa la respuesta de un repositorio con sus colaboradores.
type RespuestaRepositorio struct {
	Repositorio   *Repositorio `json:"repositorio"`
	Colaboradores []Miembro    `json:"colaboradores"`
}

// RepositorioRequest representa la estructura de una solicitud para crear un repositorio.
type RepositorioRequest struct {
	Org         string `json:"org"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Privado     bool   `json:"privado"`
}
