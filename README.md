Simple REST server that returns a given HTTP code

Usage:

```
  -code int
        Default HTTP return code (default 200)
  -l string
        listen to (default ":8080")
```

All endpoints also access the query parameter `code` that can be used to overide the default, e.g.

```
$ curl -I http://127.0.0.1:8080\?code\=200
HTTP/1.1 200 OK
Date: Fri, 14 Jun 2019 18:51:09 GMT
Content-Length: 33
Content-Type: text/plain; charset=utf-8

$ curl -I http://127.0.0.1:8080\?code\=502
HTTP/1.1 502 Bad Gateway
Date: Fri, 14 Jun 2019 18:51:13 GMT
Content-Length: 33
Content-Type: text/plain; charset=utf-8
```
