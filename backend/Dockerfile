FROM golang

EXPOSE 3000

COPY . ${GOPATH}/src/github.com/ROWA/rowa-dev/backend
WORKDIR ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/src/main

RUN ls

RUN go run server.go