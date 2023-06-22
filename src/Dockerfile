FROM golang:1.20.5 AS build

WORKDIR /app

COPY ./ ./
RUN go mod download

RUN go build -o /OpenContainerForwarder

## Deploy
FROM alpine:3.18.2

RUN apk add libc6-compat

WORKDIR /app

COPY --from=build /OpenContainerForwarder /OpenContainerForwarder

COPY templates/ /app/templates/

ENTRYPOINT ["/OpenContainerForwarder"]