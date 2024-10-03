FROM golang:1.21-alpine

EXPOSE 8080

# .env -> docker-compose -> Dockerfile forwarding
ARG APP_DIR
ARG APP_NAME
ENV APP_DIR ${APP_DIR}
ENV APP_NAME ${APP_NAME}

COPY . ${APP_DIR}
WORKDIR ${APP_DIR}
VOLUME ${APP_DIR}

# install psql
RUN apk update
RUN apk add postgresql-client
RUN apk add curl

# fetch github.com/cosmtrek/air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD air --build.cmd "go build -o ${APP_NAME} cmd/main.go" --build.bin "./${APP_NAME}"
