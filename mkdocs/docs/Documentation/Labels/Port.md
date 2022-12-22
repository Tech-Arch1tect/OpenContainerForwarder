## Synopsis

This label is used to determine which port should be proxied to within the container.

## Label

`open.container.forwarder.port`

## Value

Port, int

## Details

If the software within the container is listening on port 8080 use the following label on the container `open.container.forwarder.port=8080`

OpenContainerForwarder attempts to _guess_ the port if `open.container.forwarder.port` is not defined. The following flow is followed:

1. if `open.container.forwarder.port` is defined, this port is used.
1. if the container only exposes a single port, this port is used.
1. if there are multiple exposed ports and one of them is 80, 80 is used.
1. if there are multiple exposed ports and none of them are 80, the first exposed port found is used and a warning is issued that this is likely incorrect.
1. if there are no exposed ports, port 80 is used and issue a warning that this is likely incorrect.

It is always recommended to explicitly define `open.container.forwarder.port` to avoid confusion. This behaviour may change in the future. Defining `open.container.forwarder.port` is the only way to ensure no issues with future updates.