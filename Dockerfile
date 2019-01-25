# build stage
FROM golang:alpine AS build-env

RUN apk update \
    && apk add --no-cache git

COPY ./twitter-action.go $GOPATH/src/github.com/xorilog/twitter-action/
WORKDIR $GOPATH/src/github.com/xorilog/twitter-action/

RUN go get . \
    && go build -o twitter-action

# final stage
FROM scratch

LABEL "com.github.actions.name"="Twitter Action"
LABEL "com.github.actions.description"="Update Status (tweet) on behalf of a user"
LABEL "com.github.actions.icon"="cloud"
LABEL "com.github.actions.color"="blue"

COPY --from=build-env /src/twitter-action /go/bin/twitter-action
ENTRYPOINT ['/go/bin/twitter-action']
