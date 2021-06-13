# Go HTTP workshop

Golang http communication workshop

Requisitos:
- [go](https://golang.org/dl)
- [node](https://nodejs.org/en)
- [jq](https://stedolan.github.io/jq/download)
- [curl](https://curl.se/download.html)

For [brew](https://brew.sh/) users:
```bash
brew install go
```

```bash
brew install node
```

```bash
brew install jq
```

```bash
brew install curl
```

[Json Server](https://github.com/typicode/json-server)

```bash
npm install -g json-server
```

### Ejecutar json server

```bash
json-server --watch db.json
```

### Prueba con CURL

```bash
curl http://localhost:3000/posts | jq
```

### Ejecutar aplicacion Go

```bash
go run main.go
```