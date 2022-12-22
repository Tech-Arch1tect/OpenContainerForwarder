## Synopsis

OpenContainerForwarder configures access logs in Caddy named *hostname*.log. This label configures the logformat used. 

## Label

`open.container.forwarder.logformat`

## Value

`console` or `json` are the only valid values, string

## Details

It is planned to add new options to configure the access logs further (or disable per-hostname). e.g. to configure log rotation settings (currently Caddy defaults are used).