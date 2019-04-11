package main

import (
	"crypto/tls"
	"net/http"

	"github.com/johanbrandhorst/certify"
	"github.com/johanbrandhorst/certify/issuers/cfssl"
)

func main() {
	c := &certify.Certify{
		Issuer: &cfssl.Issuer{},
	}
	tlsConfig := &tls.Config{
		GetCertificate: c.GetCertificate,
	}

	s := http.Server{
		TLSConfig: tlsConfig,
	}

	s.ListenAndServeTLS("", "")
}
