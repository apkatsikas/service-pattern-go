FROM golang:1.20.1-buster
WORKDIR /app
COPY . .
RUN apt-get update && apt-get upgrade -y && apt-get -y install sqlite3 && go mod download
CMD make build_and_run_migrate
