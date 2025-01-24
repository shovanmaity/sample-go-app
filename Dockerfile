FROM --platform=linux/amd64 alpine:latest as certs
RUN apk --update add ca-certificates

FROM --platform=linux/amd64 registry.access.redhat.com/ubi8/ubi-minimal:8.8
COPY bin/server /usr/local/bin/server
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
CMD [ "server" ]
