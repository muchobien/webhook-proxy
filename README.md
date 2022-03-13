# Webhook Proxy

Proxy requests to tailscale network, useful for webhooks not exposed to the internet.

https://<your-domain>.dev will proxy a request to the URL specified by the upstream query parameter.

## Usage

```sh
curl -v 'https://<your-domain>.dev/wh?upstream=https://www.example.com'
```
