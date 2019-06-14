package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	listen := flag.String("l", ":8080", "listen to")
	defaultCode := flag.Int("code", 200, "Default HTTP return code")
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		codeS := r.URL.Query().Get("code")
		code, err := strconv.Atoi(codeS)
		if err != nil || code == 0 {
			code = *defaultCode
		}

		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{ \"error\": \"Roberto is awesome\" }"))
	})
	fmt.Printf("Serving at %s...\n", *listen)
	http.ListenAndServe(*listen, nil)
}
