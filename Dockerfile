FROM golang:1.16-alpine as build

# ARG GITHUB_TOKEN
# ENV ARG_GITHUB_TOKEN ${GITHUB_TOKEN}
# RUN git config --global url."https://${ARG_GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"

# ARG GITLAB_TOKEN
# ENV ARG_GITLAB_TOKEN ${GITLAB_TOKEN}
# RUN git config --global url."https://${ARG_GITLAB_TOKEN}@gitlab.com/".insteadOf "https://gitlab.com/"

WORKDIR /go/src/gitlab.com/renodesper/gokit-microservices
COPY . .

RUN rm -rf vendor .vendor*
RUN make build

# Copy into the base image
FROM gcr.io/distroless/base

# Copy the bin file
COPY --from=build /go/src/gitlab.com/renodesper/gokit-microservices/build/skeletond /skeletond

ENTRYPOINT ["/skeletond"]
EXPOSE 8000
