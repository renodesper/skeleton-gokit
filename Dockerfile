FROM golang:1.17-alpine as builder

# I use Makefile and tput, so I need to install them
RUN apk add --no-cache ncurses make

WORKDIR /go/src/gitlab.com/renodesper/gokit-microservices
COPY . .

RUN rm -rf vendor .vendor* \
  && make vendor \
  && make build

# Copy into the base image
FROM gcr.io/distroless/static:latest

# Copy the bin file
COPY --from=builder /go/src/gitlab.com/renodesper/gokit-microservices/build/skeletond /skeletond
COPY --from=builder /go/src/gitlab.com/renodesper/gokit-microservices/config/env/production.toml /production.toml

ENTRYPOINT ["/skeletond", "-config", "./production.toml"]
EXPOSE 8000
