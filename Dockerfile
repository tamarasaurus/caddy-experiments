# Build caddy with our plugin
FROM caddy:2-builder AS builder
RUN caddy-builder github.com/tamarasaurus/caddy-experiments@master

# Copy over the binary
FROM caddy:2-alpine
COPY --from=builder /usr/bin/caddy /usr/bin/caddy

WORKDIR /srv
COPY Caddyfile /srv/Caddyfile
COPY index.html /srv/index.html

EXPOSE 8080

# List the installed modules (should include gateway.middleware)
RUN caddy list-modules

CMD ["caddy", "run"]