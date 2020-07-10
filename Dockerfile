# Build caddy with our plugin
FROM caddy:2-builder AS builder
RUN caddy-builder github.com/tamarasaurus/caddy-experiments@ce3052f1a99af7f2378a6987e1cd04b0dbad8b47

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