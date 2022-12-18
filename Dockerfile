#first stage - builder
FROM golang:1.19.4-alpine3.17 as builder
WORKDIR /build
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY  . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o quote-generator

#second stage
FROM alpine:3.17.0
WORKDIR /root/
COPY --from=builder /build .
CMD ["./quote-generator"]