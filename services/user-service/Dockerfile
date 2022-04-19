# ************ STAGE - 1 : build-env *******************
FROM golang:1.18.0-alpine as build-env-user-service

#  All these steps will be cached -----
RUN mkdir /user-service
WORKDIR /user-service
# COPY go.mod and go.sum files to the workspace
COPY go.mod . 
COPY go.sum . 

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# Build the binary
RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/user-service main.go

# ********* STAGE - 2 : building minimal image (Copying only the build artifact from previous stage) ***************
FROM alpine:latest

RUN mkdir /userdata
WORKDIR /userdata

COPY --from=build-env-user-service /go/bin/user-service /go/bin/user-service

ENTRYPOINT ["/go/bin/user-service"]

