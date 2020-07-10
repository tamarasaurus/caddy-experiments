FROM caddy:2.0.0-builder AS builder
COPY . .
RUN caddy-builder github.com/tamarasaurus/caddy-experiments@master

FROM caddy:2-alpine
COPY --from=builder /usr/bin/caddy /usr/bin/caddy

WORKDIR /srv
COPY Caddyfile /srv/Caddyfile
COPY --from=build index.html /srv/index.html

EXPOSE 8080

CMD ["caddy", "run"]