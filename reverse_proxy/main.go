package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// our RerverseProxy object
type Prox struct {
	// target url of reverse proxy
	target *url.URL
	// instance of Go ReverseProxy thatwill do the job for us
	proxy *httputil.ReverseProxy
}

// small factory
func (p *Prox) New(target string) *Prox {
	url, _ := url.Parse(target)
	// you should handle error on parsing
	return &Prox{target: url, proxy: httputil.NewSingleHostReverseProxy(url)}
}

func (p *Prox) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GoProxy", "GoProxy")
	// call to magic method from ReverseProxy object
	p.proxy.ServeHTTP(w, r)
}

func main() {
	// come constants and usage helper
	const (
		defaultPort        = ":80"
		defaultPortUsage   = "default server port, ':80', ':8080'..."
		defaultTarget      = "http://127.0.0.1:8080"
		defaultTargetUsage = "default redirect url, 'http://127.0.0.1:8080'"
	)

	// flags
	port := flag.String("port", defaultPort, defaultPortUsage)
	url := flag.String("url", defaultTarget, defaultTargetUsage)

	flag.Parse()

	fmt.Println("server will run on : %s", *port)
	fmt.Println("redirecting to :%s", *url)

	// proxy
	proxy := &Prox{}
	proxy.New(*url)

	// server
	http.HandleFunc("/", proxy.handle)
	http.ListenAndServe(*port, nil)
}
