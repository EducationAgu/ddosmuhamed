FROM golang as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /deploy/server/migrations/main ./migrations/main.go &&\
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /deploy/server/main ./cmd/main.go

COPY ./migrations/ /deploy/server/migrations/
COPY ./certs/ /deploy/server/certs

FROM alpine

WORKDIR /app
COPY --from=builder ./deploy/server/ .

EXPOSE 500