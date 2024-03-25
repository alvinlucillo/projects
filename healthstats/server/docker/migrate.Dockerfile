# Dockerfile
FROM golang:1.21

WORKDIR /migrations

ARG DB_USER
ARG DB_PASSWORD
ARG DB_NAME
ARG DB_PORT
ARG DB_HOST

# Install Dockerize
ENV DOCKERIZE_VERSION v0.6.1
RUN curl -OL https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# Install migrate tool
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Install telnet and netcat
RUN apt-get update && apt-get install -y telnet netcat-openbsd

# Copy migration files and script
COPY ./database/migrations ./files
COPY ./scripts/run-migrations.sh .

RUN chmod +x run-migrations.sh

CMD ["./run-migrations.sh"]
