FROM golang:latest-alpine as cosgo
WORKDIR '/go/src/app'
COPY . .
