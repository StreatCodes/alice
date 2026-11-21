## Alice

An HTTP proxy that can upgrade requests from HTTP to HTTPS and insert headers. The proxy can be configured with a list of addresses which you can configure to inject headers. This is useful to insert an API key into HTTP requests at the proxy level. Additionally, it's useful if the environment you're working in doesn't have HTTPS available.

Here is an example config:

```json
{
	"ListenAddress": ":9571",
	"Proxies": [
        {
            "Address": "https://www.example.org",
            "Headers": {
                "ExampleKey": "Example value"
            }
        }
    ]
}
```

When sending an HTTP request to the server you must include this header `X-Alice-Address` with the address you wish to proxy the request to. E.g.

```bash
curl -H "X-Alice-Address: https://www.example.org/" http://localhost:9571
```

With the example config and request above the request will get proxied to `https://www.example.org` and the `ExampleKey: Example value` header will get injected.


### Build

This program does not use any third party dependencies.

```bash
go build
```