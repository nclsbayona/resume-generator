
########################################################
#              Build stage for dev and test            #
########################################################
FROM docker.io/golang:1.22-alpine as dev-builder
USER 35000
WORKDIR /app
COPY cmd ./cmd
COPY pkg ./pkg
COPY go.mod .
COPY go.sum .
RUN go mod tidy
RUN go test -coverprofile=coverage.out
RUN go build -o main ./cmd/main.go


########################################################
#                          Prod                        #
########################################################
FROM docker.io/alpine:3 as prod
USER 35000
WORKDIR /app
COPY --from=dev-builder /app/main .
ENTRYPOINT ["./main","-c"]
CMD ["config.yaml"]
