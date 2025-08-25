# ⚡ GitHub API

API que expone endpoints REST para consultar información de GitHub sobre organizaciones, repositorios y colaboradores.

---

## 📘 Endpoints

### 🔹 `GET /org`

Obtiene la información general de una organización.

**Parámetros:**

| Parámetro | Tipo   | Obligatorio | Descripción                          |
| --------- | ------ | ----------- | ------------------------------------ |
| `org`     | string | ✅ Sí       | Nombre de la organización en GitHub. |

**Ejemplo:**

GET /api/org?org=test

**Respuesta:**

```json
{
  "organizacion": {
    "id": 123456789,
    "login": "test-login",
    "name": "Test",
    "public_repos": 24,
    "total_private_repos": 36,
    "avatar_url": "https://avatars.githubusercontent.com/u/123456789?v=4",
    "html_url": "https://github.com/test"
  }
}
```

---

## 📦 Requisitos previos

- [Go 1.24+](https://golang.org/dl/) instalado
- Tener un token de GitHub con permisos de lectura (mínimo `read:org`)
- Conexión a internet (consume directamente la API de GitHub)

---

## 🔐 Configuración

El token se puede cargar desde una variable de entorno o archivo. Por ejemplo:

```bash
export GH_TOKEN=ghp_tu_token_personal🧑🏻‍💻
```

## 💻 Instalación

1. Clona el repositorio:

```bash
git clone <URL-del-repositorio>
```

2. Accede al directorio del proyecto:

```bash
cd <nombre-del-directorio>
```

## 🚀 Ejecución

Para ejecutar el proyecto localmente:

```bash
go run main.go
```

Acceso desde un navegador o Postman

http://localhost:8080