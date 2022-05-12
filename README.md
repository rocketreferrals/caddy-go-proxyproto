# caddy-go-proxyproto
Caddy wrapper for go-proxyproto

caddy listener wrapper `go_proxyproto` for Caddy 2 adds support for
PROXY headers using the go proxy proto listener https://github.com/pires/go-proxyproto

### JSON

Load the listener before the tls wrapper.

Via JSON config:

```json
{
  "apps": {
    "http": {
      "servers": {
        "myserver": {
          // ...
          "listener_wrappers": [
            {"wrapper": "go_proxyproto", "timeout": "5s"},
            {"wrapper": "tls"}
          ]
          // ...
        }
      }
    }
  }
}
```

Via Caddyfile config, in your global options:

```
{
  servers {
    listener_wrappers {
      go_proxyproto {
        timeout <duration>
      }
      tls
    }
  }
}
```