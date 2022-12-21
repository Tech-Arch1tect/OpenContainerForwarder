## Intro

This page demonstrates an example deploy using the [docker-compose configuration](https://github.com/Tech-Arch1tect/OpenContainerForwarder-base).

## Create docker network

This network will have OpenContainerForwarder/Caddy and any Docker containers that will be proxied to.  

```
docker network create proxy
```

## Clone the base docker-compose setup

```
cd /opt/
git clone https://github.com/Tech-Arch1tect/OpenContainerForwarder-base.git OpenContainerForwarder
```

## Copy example environment variables

```
cd /opt/OpenContainerForwarder/
cp sample.env .env
```

## Review environment variables 

- Check the VERSION variable is the latest version (or otherwise intended version)
- Unless you specifically want another TLS provider, it is advised to keep DEFAULT_TLS_PROVIDER=default
- Keep the web dashboard disabled - this is WIP/Insecure (and has no authentication currently)

## Start Open Container Forwarder & caddy

```
cd /opt/OpenContainerForwarder
docker-compose up -d
```

## Start an example container (gitea)

- Before proceeding please configure your DNS A record for your hostname to point at your server running OpenContainerforwarder (yourdomain.com will be used as an example)
- Gitea is used as an example service

```
cd /opt/
mkdir gitea
cd gitea;
vi docker-compose.yml
```

Paste in the following content (changing yourdomain.com to your domain):

```
version: "2"

networks:
  proxy:
    external: true

services:
  gitea:
    image: gitea/gitea:latest
    restart: always
    networks:
      - proxy
    volumes:
      - ./gitea:/data
    labels:
      open.container.forwarder.port: "3000"
      open.container.forwarder.hostname: "gitea.yourdomain.com"
```

Notes:

- The Gitea container is on the proxy network
- The labels define the hostname that will be configured in Caddy and the port that you want to proxy to (in this case Gitea is listening on port 3000 within the container)

### Start the gitea container

```
docker-compose up -d
```


Wait 30 seconds, and visit gitea.yourdomain.com. You should be presented with the gitea welcome page.

If you have any issues checkout the [trouble shooting](/troubleshooting) examples and create a github issue if you need any assistance.