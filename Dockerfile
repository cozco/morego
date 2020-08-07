FROM go:alpine as cosgo
WORKDIR '/go/src/app'
COPY . .
