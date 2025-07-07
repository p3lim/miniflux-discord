FROM alpine:latest AS certs

RUN apk --update add ca-certificates

FROM scratch

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --chown=65536:65536 --chmod=0777 miniflux-discord /

ENV LISTEN_PORT 8080
ENV LISTEN_ADDR 0.0.0.0

EXPOSE 8080/tcp

CMD ["/miniflux-discord"]
