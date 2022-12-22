## Synopsis

This label configures which protocol to use when forwarding to the container. e.g. if the container is listening on https then this should be configured with https. Defaults to `http`.

## Label

`open.container.forwarder.protocol`

## Value

`http` or `https` are the only valid options, string

## Details

Example usage:

http  
`open.container.forwarder.protocol=http`

https  
`open.container.forwarder.protocol=https`