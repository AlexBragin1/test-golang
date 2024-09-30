# syntax=docker/dockerfile:1

FROM golang:1.22 AS migrate

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

FROM migrate AS source

COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . ./

FROM source AS build
ARG VERSION='none'
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X gitlab.com/usdtkg/payout/config.Version=${VERSION} -X gitlab.com/usdtkg/payout/config.Build=`date -u +%Y%m%d.%H%M%S`" -o /payout-app

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

CMD ["/payout-app"]
