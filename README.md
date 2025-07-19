# Cats api

Api feita com o objetivo de aprende golang, sem se preocupar com estruta de código nem nada similar.
O objetivo puro é entender a linguagem e como ela funciona, seus tipos e sintaxe.

Para executar esse projeto use o comando:

```bash
docker compose up -d
```

O projeto iniciará na porta 8080 do localhost. :)

## Rotas

* `GET /cats` - lista todos os gatos
* `POST /cats` - cria um novo gato
* `GET /cats/:id` - retorna um gato especifico
* `PUT /cats/:id` - atualiza um gato
* `DELETE /cats/:id` - remove um gato

Para rodar os testes:

```bash
go test ./...
```

Também é possível utilizar o Makefile incluído no projeto para facilitar essas tarefas:

```bash
make run   # inicia a aplicação
make test  # executa os testes
make lint  # verifica o código
```
