FROM golang:alpine as build-stage
RUN apk add build-base
COPY . ${GOPATH}/src/github.com/ROWA/rowa-dev/backend
WORKDIR ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/src/main
RUN go build server.go


FROM alpine as production-stage

#RUN apk add tzdata
#RUN echo "Europe/Berlin" > /etc/timezone
#RUN dpkg-reconfigure -f noninteractive tzdata
EXPOSE 3000
RUN mkdir ROWA
COPY --from=build-stage /go/src/github.com/ROWA/rowa-dev/backend/src/main/ ./ROWA/

WORKDIR ${HOME}/ROWA
ENTRYPOINT [ "./server" ]
