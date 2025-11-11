FROM golang:1.25-alpine AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/server/main.go

FROM golang:1.25-alpine AS dev

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["air"]

## -----------------
## Production Stage
## -----------------
FROM alpine:latest AS prod

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/main .

RUN adduser -D appuser
USER appuser

EXPOSE 3000

CMD ["./main"]