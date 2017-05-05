package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Server stores the hostname and port number
type Server struct {
	Hostname  string // Server name
	UseHTTP   bool   // Listen on HTTP
	UseHTTPS  bool   // Listen on HTTPS
	HTTPPort  int    // HTTP port
	HTTPSPort int    // HTTPS port
	CertFile  string // HTTPS certificate
	KeyFile   string // HTTPS private key
}

// Run starts the HTTP and/or HTTPS listener
func RunServer(httpHandlers http.Handler, httpsHandlers http.Handler) {

	s := Server{
		Hostname:  "",
		UseHTTP:   true,
		UseHTTPS:  false,
		HTTPPort:  8000,
		HTTPSPort: 443,
		CertFile:  "tls/server.crt",
		KeyFile:   "tls/server.key",
	}

	port := 8000
	s.HTTPPort = port
	if s.UseHTTP && s.UseHTTPS {
		go func() {
			startHTTPS(httpsHandlers, s)
		}()

		startHTTP(httpHandlers, s)
	} else if s.UseHTTP {
		startHTTP(httpHandlers, s)
	} else if s.UseHTTPS {
		startHTTPS(httpsHandlers, s)
	} else {
		log.Println("Config file does not specify a listener to start")
	}
}

// startHTTP starts the HTTP listener
func startHTTP(handlers http.Handler, s Server) {
	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), "Running HTTP "+httpAddress(s))

	// Start the HTTP listener
	log.Fatal(http.ListenAndServe(httpAddress(s), handlers))
}

// startHTTPs starts the HTTPS listener
func startHTTPS(handlers http.Handler, s Server) {
	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), "Running HTTPS "+httpsAddress(s))

	// Start the HTTPS listener
	log.Fatal(http.ListenAndServeTLS(httpsAddress(s), s.CertFile, s.KeyFile, handlers))
}

// httpAddress returns the HTTP address
func httpAddress(s Server) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", s.HTTPPort)
}

// httpsAddress returns the HTTPS address
func httpsAddress(s Server) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", s.HTTPSPort)
}
