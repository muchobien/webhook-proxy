FROM golang:1.17.8-alpine3.15 as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o whp
# This is where one could build the application code as well.

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Copy binary to production image
COPY --from=builder /app/whp /app/whp
# Run on container startup.
CMD ["/app/whp"]
