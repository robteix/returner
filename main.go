package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	listen := flag.String("l", ":8080", "listen to")
	defaultCode := flag.Int("code", 502, "Default HTTP return code")
	errorOptions := flag.Bool("error-on-options", false, "Whether it should error on OPTIONS")
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" && !*errorOptions {
			w.WriteHeader(200)
			return
		}
		codeS := r.URL.Query().Get("code")
		code, err := strconv.Atoi(codeS)
		if err != nil || code == 0 {
			code = *defaultCode
		}

		w.WriteHeader(code)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(nginxResponse))
	})
	fmt.Printf("Serving at %s...\n", *listen)
	http.ListenAndServe(*listen, nil)
}

var nginxResponse = `<!DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="/_errors/main.css"/>
	<title>Error 502 - %{HOSTNAME}</title>
	<style>
                html{
                        background-color: #f1c40f;
                }
                body{
                        color: #fefefe;
                }
        </style>

</head>

<body>
<div class="error-middle">
        <h1>Error 502 - Bad Gateway</h1>
        <p>The 502 (Bad Gateway) status code indicates that the server, while acting as a gateway or proxy, received an invalid response from an inbound server it accessed while attempting to fulfill the request.</p>
</div>
</body>
</html>`
