## Intro

Here I'll show some example docker-compose.yml files for some common apps working with OpenContainerForwarder. These examples are not designed to be production-ready and most applications at the least require a database (not included here), but hopefully these examples will make it clear how easy it is to get an application working with OpenContainerForwarder.

## Wordpress

```
networks:
  proxy:
    external: true

services:
  wordpress:
    image: wordpress
    volumes:
      - ./wp:/var/www/html
    labels:
      open.container.forwarder.port: "80"
      open.container.forwarder.hostname: "wp.yourdomain.com"
    networks:
      - proxy
```

## Gitea

```
networks:
  proxy:
    external: true

services:
  gitea:
    image: gitea/gitea:latest
    networks:
      - proxy
    volumes:
      - ./gitea:/data
    labels:
      open.container.forwarder.port: "3000"
      open.container.forwarder.hostname: "gitea.yourdomain.com"
```

## Jenkins

```
networks:
  proxy:
    external: true

services:
  jenkins:
    image: jenkins/jenkins:lts
    volumes:
      - ./jenkins_data:/var/jenkins_home
    labels:
      open.container.forwarder.port: "8080"
      open.container.forwarder.hostname: "jenkins.yourdomain.com"
    networks:
      - proxy
```

### Nextcloud

```
networks:
  proxy:
    external: true
services:
  nextcloud:
    image: nextcloud:production-apache
    volumes:
      - "./data:/var/www/html"
    labels:
      open.container.forwarder.hostname: "nextcloud.yourdomain.com"
      open.container.forwarder.port: "80"
    networks:
      - proxy
```