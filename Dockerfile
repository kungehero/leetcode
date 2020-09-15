
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN main