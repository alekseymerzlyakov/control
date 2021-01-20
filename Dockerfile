FROM golang:1.12-alpine AS build_base
RUN apk add --no-cache git

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the Current Working Directory inside the container
WORKDIR /app
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the Go app
#RUN go build -o .control .
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o control .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o control .


FROM selenium/node-chrome:4.0.0-beta-1-prerelease-20210106
#RUN apk add ca-certificates


COPY --from=build_base /app .

#----------------------

LABEL authors=SeleniumHQ
USER root
RUN mkdir -p /Report

#--------------------

EXPOSE 8080
# Run the binary program produced by `go install`
#CMD ["/app/control"]
ENTRYPOINT ["/control"]















# 更新软件源

#
#FROM golang as builder
#
#WORKDIR /app
#COPY . .
#COPY vendor/ vendor/
#
#
#FROM robcherry/docker-chromedriver
###FROM selenium/standalone-chrome
#
#COPY --from=builder /app /usr/local/bin/webdriver_exporter
#
#RUN useradd -ms /bin/bash dev
#
#EXPOSE 9156








### /bin/sh
### version for linux
#FROM golang:1.12-alpine
##FROM robcherry/docker-chromedriver
##FROM golang:alpine3.6
#
#
#
#RUN apk add --no-cache git
#
## Set the Current Working Directory inside the container
#
#ADD . /app
#WORKDIR /app
## We want to populate the module cache based on the go.{mod,sum} files.
##RUN go mod download
#COPY go.mod .
#COPY go.sum .
#
#COPY . .
#
#
## Build the Go app
##RUN go build -o .control .
##RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o control .
#RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o control .
#
## This container exposes port 8080 to the outside world
#EXPOSE 8080
## Run the binary program produced by `go install`
#        #CMD [".control"]
#ENTRYPOINT ["/app/control"]



#FROM node:12.7.0-alpine
#WORKDIR /app
#COPY --from=builder /app/control .
## install chromedriver
#RUN apk add --update \
#    wget \
#    udev \
#    ttf-freefont \
#    chromium \
#    chromium-chromedriver \



#FROM alpine:3.12.0 AS production
#COPY --from=builder /app .
#EXPOSE 8000 40000
#ENV PORT=8000
#CMD ["/app", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "./control"]
#



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