FROM golang:latest as cosgo
WORKDIR '/go/src/app'
COPY . .
