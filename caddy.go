package caddygoproxyproto

import "github.com/caddyserver/caddy/v2"

func init() {
	caddy.RegisterModule(Wrapper{})
}

func (Wrapper) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "caddy.listeners.go_proxyproto",
		New: func() caddy.Module { return new(Wrapper) },
	}
}