package caddygoproxyproto

import (
	"net"
	"time"

	"github.com/caddyserver/caddy/v2"
	"github.com/pires/go-proxyproto"
)

type Wrapper struct {
	Timeout caddy.Duration `json:"timeout,omitempty"`
}

func (pp *Wrapper) Provision(ctx caddy.Context) error {
	return nil
}

func (pp *Wrapper) WrapListener(l net.Listener) net.Listener {
    proxyListener := &proxyproto.Listener{
        Listener: l,
		ReadHeaderTimeout: time.Duration(pp.Timeout),
	}
	return proxyListener
}
