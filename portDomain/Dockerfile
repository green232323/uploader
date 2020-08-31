# running app stage
FROM alpine:3.12

RUN apk add --no-cache ca-certificates

WORKDIR /usr/src/app

COPY portDomain .

ENTRYPOINT ["./portDomain"]
USER guest
