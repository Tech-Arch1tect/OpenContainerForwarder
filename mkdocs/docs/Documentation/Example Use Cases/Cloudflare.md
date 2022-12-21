## Intro

If all / the majority of your websites are behind a  reverse proxy/cloudflare there are some useful default's that you can set globally.


## Using Cloudflare DNS for the TLS provider

- Certificates are verified through DNS records, caddy/cloudflare DNS plugin automatically add the required DNS records for verification using the Cloudflare API
- You need a Cloudflare API key which has permissions to add DNS records for your domains

Within `/opt/OpenContainerForwarder/.env` configure the TLS provider + Cloudflare API key:

```
DEFAULT_TLS_PROVIDER=cloudflare
CLOUDFLARE_API_KEY=xxxxxx
```

With this configuration all containers will use Cloudflare DNS for SSL verification using the specified API key, unless the container overrides these settings using the open.container.forwarder labels (open.container.forwarder.tlsprovider and open.container.forwarder.cloudflareapiKey)


## Trusted Proxies

Caddy by default doesn't trust any x-forwarded-* etc headers sent to it (e.g. sent by the Cloudflare proxy). By defining trusted proxies we can allow caddy to forward the x-forwarded-* headers down to the container backends when the requests are from Cloudflare IP's.

For Cloudflare you can grab the IP ranges from [here](https://www.cloudflare.com/en-gb/ips/) and then configure them as your trusted proxies default with the following .env entry:

```
DEFAULT_TRUSTED_PROXIES=xxx.xxx.xxx.xxx/XX x.x.x.x/XX xxx.x.xxx.x/XX
```

This can be overridden on a per-container basis using the container label `open.container.forwarder.trustedproxies` (e.g. if a container/hostname is not using cloudflare)