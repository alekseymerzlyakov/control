# /bin/sh
# version for linux
FROM golang:1.12-alpine
#FROM golang:1.13.10-alpine3.11
RUN apk add --no-cache git
# Set the Current Working Directory inside the container
WORKDIR /app
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
# Build the Go app
#RUN go build -o .control .
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o control .
#RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o control .
# This container exposes port 8080 to the outside world
EXPOSE 8080
# Run the binary program produced by `go install`
        #CMD [".control"]
ENTRYPOINT ["/app/control"]


#FROM scratch
#ADD . ./
#ENTRYPOINT ["/control"]



## work version for MAC
#FROM golang:1.12-alpine
#RUN apk add --no-cache git
## Set the Current Working Directory inside the container
#WORKDIR /app
## We want to populate the module cache based on the go.{mod,sum} files.
#COPY go.mod .
#COPY go.sum .
#RUN go mod download
#COPY . .
## Build the Go app
#RUN go build -o ./control .
###RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app .
## This container exposes port 8080 to the outside world
#EXPOSE 8080
## Run the binary program produced by `go install`
#CMD ["./control"]