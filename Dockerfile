FROM golang:1.20.1-buster
WORKDIR /app
COPY . .
RUN apt-get update && apt-get upgrade && apt-get install sqlite3 && go mod download \
    && echo '#!/bin/bash\ngo "$@"' > /usr/bin/go1.20.1 && chmod +x /usr/bin/go1.20.1
CMD make setup-db && make run
