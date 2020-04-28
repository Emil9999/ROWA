FROM golang:alpine as build-stage
RUN apk add build-base
COPY . ${GOPATH}/src/github.com/ROWA/rowa-dev/backend
WORKDIR ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/src/main
RUN go build server.go


FROM alpine as production-stage

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
EXPOSE 3000
RUN mkdir ROWA
COPY --from=build-stage ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/src/main/ ./ROWA/
COPY --from=build-stage ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/keys/ ./ROWA/
RUN cd ROWA && update-ca-certificates

WORKDIR ${HOME}/ROWA
ENTRYPOINT [ "./server" ]
