FROM golang as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /deploy/server/migrations/main ./migrations/main.go &&\
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /deploy/server/main ./cmd/main.go

COPY ./migrations/ /deploy/server/migrations/

FROM alpine

WORKDIR /app
COPY ./backend.ovpn /etc/openvpn/backend.ovpn
COPY --from=builder ./deploy/server/ .

EXPOSE 500