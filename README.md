# Step by step

Setup the project and install gokit packages:

    ```sh
    go mod init gitlab.com/renodesper/gokit-microservices
    go get github.com/go-kit/kit
    ```

## Creating a service (ex: health check)

1. Create service contract.
2. Implement the service.
3. Create endpoint contract.
4. Implement the endpoint. Do not forget to use the service that we created before inside the endpoint.
5. Create decoder for the health check service.
6. Implement the endpoint as a route in the transport.

## Run the app

Run the main go file:

    ```sh
    go run cmd/main.go
    ```

Use Makefile:

    ```sh
    make run
    ```
