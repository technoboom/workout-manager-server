# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.9

# install dependencies
RUN go get github.com/gorilla/mux

# copy the local package files to the container workspace
ADD . /go/src/github.com/technoboom/workout-manager-server/users

# setting up working directory
WORKDIR /go/src/github.com/technoboom/workout-manager-server/users

# build the users service inside the container
RUN go install github.com/technoboom/workout-manager-server/users

# run the users microservice when the container starts
ENTRYPOINT /go/bin/users

# service listens on port 8080
EXPOSE 8080
