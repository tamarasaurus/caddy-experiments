package gatewaymiddleware

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(GatewayMiddleware{})
	httpcaddyfile.RegisterHandlerDirective("gateway_middleware", parseCaddyfile)
}

type GatewayMiddleware struct {
}

func (GatewayMiddleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.gateway_middleware",
		New: func() caddy.Module { return new(GatewayMiddleware) },
	}
}

func (s *GatewayMiddleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	return nil
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	t := new(GatewayMiddleware)
	err := t.UnmarshalCaddyfile(h.Dispenser)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s GatewayMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, _ caddyhttp.Handler) error {
	// w.WriteHeader(http.)
	w.Write([]byte("Admin token"))

	return nil
}

var (
	_ caddyhttp.MiddlewareHandler = (*GatewayMiddleware)(nil)
	_ caddyfile.Unmarshaler       = (*GatewayMiddleware)(nil)
)