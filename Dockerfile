FROM index.docker.io/library/golang:alpine as build
WORKDIR /src
COPY main.go go.mod .
RUN go build -o helloserver -ldflags="-s -w"

FROM index.docker.io/library/alpine:edge
LABEL maintainer "Ramon van Stijn <ramons@nl.ibm.com>"
RUN addgroup -g 1970 hello \
    && adduser -u 1970 -G hello -s /bin/sh -D hello
COPY --chown=hello:hello --from=build /src/helloserver /app/helloserver
USER hello
EXPOSE 1970
ENTRYPOINT /app/helloserver
