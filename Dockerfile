#first stage - builder
FROM golang:1.19.4-alpine3.17 as builder
COPY . /build
WORKDIR /build
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o quote-generator

#second stage
FROM alpine:3.17.0
WORKDIR /root/
COPY --from=builder /build .
CMD ["./quote-generator"]