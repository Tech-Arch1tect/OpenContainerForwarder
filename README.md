# OpenContainerForwarder

# This was written for learning, it is not a production ready software. Bug will be common, security may not have been considered for every scenario, etc.

## Summary

OpenContainerForwarder automates the process of reverse proxying to containers.

Currently only Docker containers are supported and only Caddy can be used as a reverse proxy. However it may be possible to expand this support to other container engines + web servers / proxies in the future.

Limitations:

- Only containers running on the same machine as OpenContainerForwarder can be proxied to.

How:

- OpenContainerForwarder uses the Docker API to get information about running containers. 
- Using container labels OpenContainerForwarder identifies which containers to proxy to, using which hostname etc.
- OpenContainerForwarder then generates a Caddy configuration and using the Caddy admin API, loads this configuration.


Priorities:

- Code cleanup - The code is currently not very pretty or particularly efficient.
- Code quality - The code is not of the best quality. It is not commented and legibility is low.
- Documentation - This is being worked on slowly but surely. When available this will be updated.
- Features - New features are currently low priority while code cleanup and quality are focussed on.
