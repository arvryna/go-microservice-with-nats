# ************ STAGE - 1 : build-env *******************
FROM golang:1.18.0-alpine as build-env-tran-service

#  All these steps will be cached -----
RUN mkdir /tran-service
WORKDIR /tran-service
# COPY go.mod and go.sum files to the workspace
COPY go.mod . 
COPY go.sum . 

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# Build the binary
RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/tran-service main.go

# ********* STAGE - 2 : building minimal image (Copying only the build artifact from previous stage) ***************
FROM alpine:latest

RUN mkdir /transdata
WORKDIR /transdata

COPY --from=build-env-tran-service /go/bin/tran-service /go/bin/tran-service

ENTRYPOINT ["/go/bin/tran-service"]
