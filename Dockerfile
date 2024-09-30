# syntax=docker/dockerfile:1

FROM golang:1.22 AS migrate

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . ./


RUN CGO_ENABLED=0 GOOS=linux go build -o /test-golang


EXPOSE 8080

CMD ["/test-golang"]
