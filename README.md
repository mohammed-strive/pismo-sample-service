# Pismo Service
This is an API service written in Go, with a Postgres backend. The [pismo service file](pismo_service.yml) has the service description. There is a [docker compose file](docker_compose.yml) to instantiate and run the service. There are migration files that will run as part of the service start. The service is written as a typical api service in go with controllers, services and repository layers.

## Prerequisites
1. Golang [link](https://go.dev)
2. Docker and Docker compose [link](https://docs.docker.com/get-started/)

## Installation
### Local
1. Clone the repo.
2. In the project root run `go mod download`. This will install the required dependencies.
3. Setup the DB credentials in the terminal. These are part of the docker compose file. `DB_HOST`, `DB_USER`, `DB_PORT`, `DB_PASSWORD`, `DB_NAME`.
4. Run the db service - `docker-compose up db`. This will setup the database.
5. Run the migrate service - `docker-compose up migrate`. This will run the DB migrations which are part of the repo.
6. Run `go run ./...`. This will run the gin server locally on port `8082`.

### Containerized
1. Run `docker-compose up`. All the necessary services are start and the api server runs at `8082`.

## Notes
1. Run tests - `go test ./...`. Tests are written using `Ginkgo` and `Gomega`.
2. Migrations are written using `go-migrate`. [link](https://github.com/golang-migrate/migrate)