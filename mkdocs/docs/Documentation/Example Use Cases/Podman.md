### Warning

This section of the documentation is currently WIP

## Intro

Podman is not currently supported however has been tested and appears to be working. This is due to the podman API implementing a Docker compatible API.

## The testing environment

- Podman running in root mode (So it can bind to port 80/443 - this should be able to be resolved root-less using net.ipv4.ip_unprivileged_port_start as described [here](https://github.com/containers/podman/blob/main/rootless.md))
- Podman socket available at /var/run/podman/podman.sock
- Arch used as host OS
- Podman using netavark as network backend with aardvark-dns setup for container DNS resolution

## Instructions

1. Set the DOCKER_HOST var so docker & docker-compose commands use the podman socket
```
export DOCKER_HOST='unix:///var/run/podman/podman.sock'
```
1. Create the proxy network
```
podman network create proxy
```
1. (optional) For dev only, disable buildkit
```
export DOCKER_BUILDKIT=0
```
1. Follow the [Quick start](../../../Quick Start/). Ignoring the "Create docker network" step.
1. Within the .env file set the CONTAINER_SOCKET variable to the podman socket location.