## Synopsis

This label configures the SNI hostname that will be configured in Caddy. 

Caddy will use this/these hostnames for the SSL certificate.

## Label

`open.container.forwarder.hostname`  
`open.container.forwarder.hostname.*`

## Value

Valid hostname, `string`

## Details

If only a single hostname is desired simply configure this with `open.container.forwarder.hostname=my.hostname.com`.

If multiple hostnames are desired to proxy to the same container use the wildcard. e.g. to use one.hostname.com, two.hostname.com and three.hostname.com use the following labels:

```
open.container.forwarder.hostname.1=one.hostname.com
open.container.forwarder.hostname.2=two.hostname.com
open.container.forwarder.hostname.3=three.hostname.com
```

Each of these hostnames will be configured in Caddy.