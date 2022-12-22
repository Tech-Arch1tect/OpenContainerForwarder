## Synopsis

This label configures how SSL certificates are obtained.

## Label

`open.container.forwarder.tlsprovider`

## Value

`default`, `cloudflare` or `internal` are the only valid options, string

## Details

- Default - This uses Caddy's default http validation.
- Cloudflare - This configures Caddy to use DNS validation, using Cloudflare DNS and API.
- Internal - This configures Caddy to generate it's own self-signed certificates.