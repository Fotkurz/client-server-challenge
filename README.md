# client-server-challenge

Client - Server challenge for fullcycle course.

## Execução

Basta executar o arquivo `main.go` presente nas pastas `cmd` de cada projeto (client e server).

Primeiro execute o servidor:

```bash
cd server
go run cmd/main.go
```

Em outro terminal, execute o client:

```bash
cd client
go run cmd/main.go
```

PS: Mantive o timeout do banco em 20 ms porque com 10 ms não estava conseguindo inserir nenhum registro. Também adicionei um endpoint para
listar todos os registros do banco (`curl localhost:8080/all`).
