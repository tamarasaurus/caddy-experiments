package gateway/middleware

import "github.com/caddyserver/caddy/v2"

func init() {
	caddy.RegisterModule(GatewayMiddleware{})
}

// GatewayMiddleware is an example; put your own type here.
type GatewayMiddleware struct {
}

// CaddyModule returns the Caddy module information.
func (GatewayMiddleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "gateway.middleware",
		New: func() caddy.Module { return new(GatewayMiddleware) },
	}
}