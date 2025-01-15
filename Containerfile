# build
FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY . .

ARG CGO_ENABLED
ARG GOOS
ARG GOARCH

RUN go mod download && CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -a -installsuffix cgo -o gin-album-app .

# Deploy
FROM golang:1.23.2-alpine

WORKDIR /

COPY --from=builder /app/gin-album-app .

ENTRYPOINT ["./gin-album-app"]