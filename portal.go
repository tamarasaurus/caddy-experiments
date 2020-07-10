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

// CaddyModule Module registration
func (GatewayMiddleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.gateway_middleware",
		New: func() caddy.Module { return new(GatewayMiddleware) },
	}
}

// UnmarshalCaddyfile No config
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
	w.Write([]byte("Return it as JSON"))
	// Read the POST body
	// Perform the HTTP request
	// Check the response
	// Return the response as JSON
	return nil
}

var (
	_ caddyhttp.MiddlewareHandler = (*GatewayMiddleware)(nil)
	_ caddyfile.Unmarshaler       = (*GatewayMiddleware)(nil)
)