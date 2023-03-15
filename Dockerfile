FROM golang:1.20.1-buster
WORKDIR /app
RUN apt-get update && apt-get upgrade && apt-get install sqlite3 && go mod download
CMD make build_and_run_migrate
