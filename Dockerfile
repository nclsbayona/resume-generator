
########################################################
#              Build stage for dev and test            #
########################################################
FROM docker.io/golang:1.22-alpine as dev-builder
WORKDIR /app
COPY cmd ./cmd
COPY domain ./domain
COPY ports ./ports
COPY usecase ./usecase
COPY adapters ./adapters
COPY go.mod .
COPY go.sum .
RUN go mod tidy
RUN go test ./domain/*.go -coverprofile=coverage-domain.out
RUN go test ./ports/*.go -coverprofile=coverage-ports.out
RUN go test ./usecase/*.go -coverprofile=coverage-usecase.out
RUN go build -o main ./cmd/main.go


########################################################
#                          Prod                        #
########################################################
FROM docker.io/alpine:3 as prod
USER 35000
WORKDIR /app
COPY --from=dev-builder /app/main .
ENTRYPOINT ["./main","-c","config","-i","yaml","-o","html"]
CMD ["config.yaml"]
# Keep in mind you would need to add your templates to the running container