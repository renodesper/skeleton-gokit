# Summary

This project was created to accommodates the needs of a bootstrap structure for my future project.

## Todo

- [ ] Add testing
- [ ] Add more testing 😂
- [ ] Add new type of logging (ex: log to ELK)
- [ ] Add system monitoring and alerting toolkit implementation (prometheus)
- [ ] Add circuit breaker implementation
- [ ] Add rate limiter implementation
- [ ] Add service discovery implementation (consul, etcd, zookeeper, etc)
- [ ] 2FA

## Setup database

Make sure the postgres has been up and running. If it is not, we can run it in the docker:

```sh
docker run --name postgres-svc -p 5432:5432 \
    -e POSTGRES_DB=project \
    -e POSTGRES_USER=user \
    -e POSTGRES_PASSWORD=password \
    -v ~/Tmp/postgres_data:/var/lib/postgresql/data \
    -d postgres:13.2-alpine
```

or

```sh
docker start <container_id>
```

Install `golang-migrate`:

```sh
brew install golang-migrate
```

Migrate the current schema:

```sh
export POSTGRESQL_URL='postgres://user:password@127.0.0.1:5432/project?sslmode=disable'
migrate -database ${POSTGRESQL_URL} -path config/db/migrations up
```

Common error:

```sh
error: Dirty database version 2. Fix and force version.
```

Solution:

```sh
migrate -database ${POSTGRESQL_URL} -path config/db/migrations force <version - 1>
```

## Creating a service (ex: health check)

1. Create service contract.
2. Implement the service.
3. Create endpoint contract.
4. Implement the endpoint. Do not forget to use the service that we created before inside the endpoint.
5. Create decoder for the health check service.
6. Implement the endpoint as a route in the transport.

## Run the app

Use Makefile with hot reload enabled ([air](https://github.com/cosmtrek/air) will automatically installed):

```sh
make watch
```

Use Makefile with hot reload disabled:

```sh
make run
```

Run the main go file:

```sh
go run cmd/main.go
```
