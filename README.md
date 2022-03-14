# caddy-go-proxyproto
Caddy wrapper for go-proxyproto

caddy listener wrapper `go_proxyproto` for Caddy 2 adds support for
PROXY headers using the go proxy proto listener https://github.com/pires/go-proxyproto

### JSON

This module is only designed to work with JSON configurations in Caddy.
Load the listener before the tls wrapper

```js
{
  "apps": {
    "http": {
      "servers": {
        "myserver": {
          // ...
          "listener_wrappers":[{"wrapper": "go_proxyproto", "timeout": "5s"}, {"wrapper":"tls"}]
          // ...
        }
      }
    }
  }
}
```
