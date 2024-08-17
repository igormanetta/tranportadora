FROM golang:1.22.3-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /app/transportadora ./cmd/transportadora/main.go

FROM ubuntu:22.04

WORKDIR /app

COPY --from=builder /app/transportadora /app/transportadora

RUN mkdir -p infra/pgstore/migrations
RUN mkdir -p controller/swagger

COPY ./infra/pgstore/migrations/. ./infra/pgstore/migrations/
COPY ./controller/swagger/. ./controller/swagger

CMD ["/app/transportadora"]