FROM golang:1.17.8-alpine3.15 as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o webhook-proxy
# This is where one could build the application code as well.


FROM alpine:latest as tailscale
WORKDIR /app
ENV TSFILE=tailscale_1.22.1_amd64.tgz
RUN wget https://pkgs.tailscale.com/stable/${TSFILE} && \
  tar xzf ${TSFILE} --strip-components=1

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Copy binary to production image
COPY --from=builder /app/webhook-proxy /app/webhook-proxy
COPY --from=builder /app/start.sh /app/start.sh
COPY --from=tailscale /app/tailscaled /app/tailscaled
COPY --from=tailscale /app/tailscale /app/tailscale
RUN mkdir -p /var/run/tailscale /var/cache/tailscale /var/lib/tailscale

# Run on container startup.
CMD ["/app/start.sh"]
