FROM golang:1.13.4
COPY . /Users/eleonoralester/go/src/github.com/elstr/example-services
# COPY . /go/src/github.com/elstr/example-services
WORKDIR /Users/eleonoralester/go/src/github.com/elstr/example-services
# WORKDIR /go/src/github.com/elstr/example-services
# RUN go install ./cmd/...
RUN go install -ldflags="-s -w" 