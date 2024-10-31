package main

import (
	"flag"
	"github.com/YuanJey/debug-server/internal/websockifygo"
	"log"
	"strings"
)

func main() {
	var keyPem string
	var certPem string
	var url string
	var port string
	var debugAddr string

	flag.StringVar(&port, "port", "0.0.0.0:9000", "SSL key.pem")
	flag.StringVar(&debugAddr, "addr", "127.0.0.1:2345", "SSL key.pem")
	flag.StringVar(&keyPem, "key", "", "SSL key.pem")
	flag.StringVar(&certPem, "cert", "", "SSL cert.pem")
	flag.StringVar(&url, "url", "/debug", "url path to proxy, e.g. /vnc")

	flag.Parse()

	if !strings.HasPrefix(url, "/") {
		url = "/" + url
	}

	wsp := &websockifygo.WSproxy{
		URL:     url,
		Target:  debugAddr,
		KeyPem:  keyPem,
		CertPem: certPem,
	}
	log.Println("Proxying", url, "to", debugAddr, "on", port)
	wsp.Serve(port)
}
