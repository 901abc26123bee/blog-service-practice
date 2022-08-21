FROM golang:1.18.1
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .
RUN         go mod download
RUN         go build -o app
EXPOSE 8080
ENTRYPOINT  ["./app"]