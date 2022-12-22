## Synopsis

This label can be used to restrict access to a website to a CIDR range of IP's.

## Label

`open.container.forwarder.restrictip`
`open.container.forwarder.restrictip.*`

## Value

CIDR range, string

## Details

Example usage:

Single CIDR range:  
`open.container.forwarder.restrictip=x.x.x.x/XX`

Multiple ranges:
`open.container.forwarder.restrictip.1=x.x.x.x/XX`
`open.container.forwarder.restrictip.2=x.x.x.x/XX`
`open.container.forwarder.restrictip.3=x.x.x.x/XX`

### Note

If the website is behind another reverse proxy you will need to configure [trusted proxies](../Trusted Proxies) as well. Configuring a trusted proxy allows Caddy to trust the x-forwarded-for header sent from the proxy which will then be used to determine the visitors IP address.