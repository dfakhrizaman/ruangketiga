FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/server
RUN go build -o /go/bin/ruangketiga

FROM alpine:3.17

COPY --from=build /go/bin/ruangketiga /ruangketiga

EXPOSE 8080

ENTRYPOINT ["/ruangketiga"]
