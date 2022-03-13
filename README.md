# Webhook Proxy

Proxy requests to tailscale network using gcloud run, useful for webhooks not exposed to the internet.

https://<your-domain>.dev will proxy a request to the URL specified by the service query parameter.

## Usage

```sh
curl -v 'https://<your-domain>.dev/wh?service=https://www.example.com'
```

## Deploy

```sh
gcloud run deploy
```
