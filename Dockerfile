FROM golang:alpine
RUN apk add git

# install dep
RUN go get github.com/golang/dep/cmd/dep

# create a working directory
WORKDIR /go/src/go-tourney

# add Gopkg.toml and Gopkg.lock
ADD Gopkg.toml Gopkg.toml
ADD Gopkg.lock Gopkg.lock

# install packages
# --vendor-only is used to restrict dep from scanning source code
# and finding dependencies
RUN dep ensure --vendor-only

# add source code
ADD . /go/src/go-tourney

EXPOSE 8080 

# run main.go
CMD ["go", "run", "main.go"]
