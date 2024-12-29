FROM golang:1.24rc1-alpine3.21 as builder

WORKDIR /app

RUN apk add --no-cache git bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o tmp/app ./cmd/main.go

FROM alpine:3.21

WORKDIR /app
COPY --from=builder /app/tmp/app .

EXPOSE 8080
CMD ["./app"]