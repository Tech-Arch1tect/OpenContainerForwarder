## Synopsis

This label is used to trust a proxy sending x-forwarded-* headers. Without configuring Trusted Proxies the x-forwarded-for headers will be ignored by Caddy. Read more about Caddy behaviour [here](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy#trusted_proxies)

## Label

`open.container.forwarder.trustedproxies`

## Value

Space seperated list of CIDR ranges, string

## Details

This label should be used when behind another reverse_proxy e.g. Cloudflare.

Using Cloudflare as an example use-case, we can take the IP ranges from [here](https://www.cloudflare.com/en-gb/ips/) and create a label like `open.container.forwarder.trustedproxies= x.x.x.x/XX x.x.x.x/XX x.x.x.x/XX` etc.

With trusted proxies configured, Caddy will trust that the configured CIDR ranegs will send good x-forwarded-* values.