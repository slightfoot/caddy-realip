Middleware for restoring real ip information when running caddy behind a proxy. Will allow other middlewares to simply use `r.RemoteAddr` instead of decoding `X-Forwarded-For` themselves.

Analogous to nginx's [realip_module](http://nginx.org/en/docs/http/ngx_http_realip_module.html)

Checks whitelist of authorized proxy servers so we don't arbitrarily trust headers from anybody.
The real IP module is useful in situations where you are running caddy behind a proxy server. In these situations, the actual client IP will be stored in an HTTP header, usually  X-Forwarded-For. The problem this creates is that other directives that rely on the client IP address, like ipfilter or git, will not always work properly in these scenarios.
This middleware will seamlessly and securely read the real IP address from the appropriate header and replace the proxied IP in the request with the real one.
Syntax

```
realip {
    header name
    from   cidr
    strict
}
```

name is the name of the header containing the actual IP address. Default is  X-Forwarded-For.
cidr is the address range of expected proxy servers. As a security measure, IP headers are only accepted from known proxy servers. Must be a valid cidr block notation. This may be specified multiple times.
strict, if specified, will reject requests from unkown proxy IPs with a 403 status. If not specified, it will simply leave the original IP in place.
CIDR blocks

CIDR is a standard notation for specifying IP ranges. To allow a single IP, use /32 after the ip to specify no mask: 123.222.31.4/32. To allow all IPs (and accept an  X-Forwarded-For header from anybody), use 0.0.0.0/0. Most cloud services should have their ip ranges published in this format somewhere.
Presets

There is a helper supplied if you want to use Caddy behind Cloudflare. Simply specify  `realip cloudflare` in your caddyfile to activate it using a built-in IP list.
Additional presets would be welcome by pull request for other cloud providers.

### Example

Simple usage to read X-Forwarded-For from a few known IPs only:

```
realip {
    from 1.2.3.4/32
    from 2.3.4.5/32
}
```
