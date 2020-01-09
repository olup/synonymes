############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN go build -o /go/bin/synonyms
############################
# STEP 2 build a small image
############################
FROM alpine

WORKDIR /go/bin/
RUN mkdir ./db
# Copy our static executable.
COPY --from=builder /go/bin/synonyms ./synonyms
RUN ls

CMD ["./synonyms"]