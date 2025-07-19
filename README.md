# Cats API

This project provides a small REST service for managing cats. It was created as
a learning exercise to showcase how to build an API in Go using the Gin
framework. The service keeps data in memory and is not intended for production
use.

## Running with Docker

Launch the API with Docker Compose:

```bash
docker compose up -d
```

Once the container starts, the API will be available at
<http://localhost:8080>.

You can also run the application directly on your machine:

```bash
go run ./cmd/cats_api
```

## Routes

* `GET /cats` – list all cats
* `POST /cats` – create a new cat
* `GET /cats/:id` – retrieve a cat by ID
* `PUT /cats/:id` – update a cat
* `DELETE /cats/:id` – remove a cat

## Tests

Run the unit tests with:

```bash
go test ./...
```

Também é possível utilizar o Makefile incluído no projeto para facilitar essas tarefas:

```bash
make run   # inicia a aplicação
make test  # executa os testes
make lint  # verifica o código
```
