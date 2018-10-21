FROM golang:1.9.2-alpine3.7

RUN mkdir -p /go/src/github.com/panicpanicpanic/filament
WORKDIR /go/src/github.com/panicpanicpanic/filament
ADD . /go/src/github.com/panicpanicpanic/filament

# Install dep and dependancies
RUN apk --no-cache add curl git && \
    curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && \
    chmod +x /usr/local/bin/dep
RUN dep ensure -vendor-only
