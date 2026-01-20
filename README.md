# API SocialMeli (Back-end)

API REST desenvolvida em Go(Golang), construída no contexto de estudos/bootcamp (Mercado Livre) para praticar **arquitetura em camadas**, **persistência com PostgreSQL (GORM)**, **documentação com Swagger** e **testes automatizados**.


## 🧩 Stack Tecnológica

- **Go**: 1.24
- **HTTP**: Gin (`github.com/gin-gonic/gin`)
- **ORM/DB**: GORM + PostgreSQL
- **Docs**: Swagger via `swaggo` (`gin-swagger`)
- **Testes**: `testing` + `testify`
- **Infra**: Docker + Docker Compose


## 🏗️ Arquitetura e organização

Estrutura (alto nível) seguindo uma separação entre **entrada da aplicação**, **camada de aplicação**, **domínio** e **infraestrutura**:

- `cmd/api/`: ponto de entrada da API (`main.go`)
- `internal/application/`: serviços de aplicação e contratos (interfaces)
- `internal/domain/`: modelos e DTOs (contratos de request/response)
- `internal/infra/`: detalhes de infraestrutura (DB, HTTP/controllers, routes, repositórios)
- `docs/`: artefatos Swagger gerados (`swagger.yaml/json`, `docs.go`)

> A aplicação executa `AutoMigrate` no startup para criar/atualizar tabelas no Postgres.


## 📦 Funcionalidades (User Stories)

- ✓ Follow / Unfollow de usuários
- ✓ Listar seguidores
- ✓ Listar usuários seguidos
- ✓ Publicação de posts
- ✓ Feed das últimas 2 semanas
- ✓ Ordenação por data
- ✓ Ordenação alfabética
- ✓ Posts promocionais
- ✓ Contagem de produtos em promoção


## 🧠 Como executar

### 🐳 Opção recomendada: Docker (API + Postgres)

1) Crie um arquivo `.env` na raiz do projeto (ele é **ignorado** pelo Git).

Exemplo:

```env
POSTGRES_HOST=db
POSTGRES_PORT=5432
POSTGRES_USER=teste
POSTGRES_PASSWORD=teste
POSTGRES_DB=socialmeli
```


2) Suba os serviços:

```bash
docker compose up --build
```

- **API**: `http://localhost:8080`
- **Postgres**: `localhost:5432` (mapeado pelo Compose)

Para desligar/remover volumes:

```bash
docker compose down -v
```

### 📍 Opção local (Go na máquina + Postgres via Docker)

1) Suba apenas o banco:

```bash
docker compose up -d db
```

2) Ajuste seu `.env` para apontar para `localhost`:

```env
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=socialmeli
```


3) Rode a API:

```bash
go run ./cmd/api
```

## 📄 Documentação (Swagger)

A API expõe Swagger UI em:

- [Swagger UI](http://localhost:8080/swagger/index.html)

E os artefatos gerados ficam em `docs/swagger.yaml` e `docs/swagger.json`.

### Regenerar Swagger (quando mudar anotações)

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g cmd/api/main.go -o docs
```

> Dica: se perceber divergência entre rotas reais e Swagger, regenere os artefatos com o comando acima.

## ✉️ Rotas (resumo)

Base path: **`/api`**

- **Users**
  - `POST /api/users`
  - `GET /api/users`
  - `GET /api/users/:user_id`
- **Products**
  - `POST /api/products`
  - `GET /api/products`
  - `GET /api/products/:product_id`
- **Posts**
  - `POST /api/posts`
  - `GET /api/posts/users/:user_id`
  - `GET /api/posts/promo`
  - `GET /api/posts/promo/count`
- **Follow**
  - `POST /api/users/follow`
  - `POST /api/users/unfollow`
  - `GET /api/users/:user_id/followers/count`
  - `GET /api/users/:user_id/followers/list`
  - `GET /api/users/:user_id/followed/list`

Para payloads/validações e exemplos, use o **Swagger** como contrato principal.

## 🧪 Testes

Rodar todos os testes:

```bash
go test ./...
```

Com cobertura:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

## Pronto para acoplar um Front-end

- **Contrato**: o Swagger serve como base para integração (inclusive para geração de client).
- **JSON + HTTP**: endpoints REST sob `/api`.
- **Ambientes**: o projeto já usa `.env`;