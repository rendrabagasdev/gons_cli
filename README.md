# Gons CLI

A command-line tool for scaffolding [Gons](https://github.com/rendrabagasdev/gons) framework projects. Quickly create new projects and generate boilerplate files for models, controllers, services, requests, and middlewares.

## Requirements

- [Go](https://golang.org/dl/) 1.21+
- [Git](https://git-scm.com/) (required for `new` command)

## Installation

```bash
go install github.com/rendrabagasdev/gons_cli@latest
```

Or build from source:

```bash
git clone https://github.com/rendrabagasdev/gons_cli.git
cd gons_cli
go build -o gons .
```

## Usage

```
gons [command] [arguments]
```

## Commands

### `new` — Create a new project

Clones the Gons framework template and sets up a new project directory.

```bash
gons new <project-name>
```

**Example:**

```bash
gons new my-app
cd my-app
go mod tidy
```

---

### `make:model` — Generate a model

Creates a new model file inside `app/models/`.

```bash
gons make:model <name>
```

**Example:**

```bash
gons make:model user
# Creates: app/models/User.go
```

---

### `make:controller` — Generate a controller

Creates a new controller file inside `app/http/controllers/`. The `Controller` suffix is added automatically if not provided.

```bash
gons make:controller <name>
```

**Example:**

```bash
gons make:controller user
# Creates: app/http/controllers/UserController.go
```

---

### `make:service` — Generate a service

Creates a new service file inside `app/http/services/`. The `Service` suffix is added automatically if not provided.

```bash
gons make:service <name>
```

**Example:**

```bash
gons make:service user
# Creates: app/http/services/UserService.go
```

---

### `make:request` — Generate a request

Creates a new request file inside `app/http/requests/`. The `Request` suffix is added automatically if not provided.

```bash
gons make:request <name>
```

**Example:**

```bash
gons make:request store-user
# Creates: app/http/requests/store_user_request.go
```

---

### `make:middleware` — Generate a middleware

Creates a new middleware file inside `app/http/middlewares/`. The `Middleware` suffix is added automatically if not provided.

```bash
gons make:middleware <name>
```

**Example:**

```bash
gons make:middleware auth
# Creates: app/http/middlewares/auth_middleware.go
```

---

## Notes

- The `make:*` commands must be run from the **root directory** of a Gons project (where the expected `app/` directory structure exists).
- If the target file already exists, the command will abort without overwriting it.

## License

Redistributable licenses place minimal restrictions on how software can be used, modified, and redistributed.

MIT [LICENSE](LICENSE)
