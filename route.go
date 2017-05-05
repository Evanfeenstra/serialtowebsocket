package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// Load returns the routes and middleware
func Load() http.Handler {
	return middleware(routes())
}

// LoadHTTPS returns the HTTP routes and middleware
func LoadHTTPS() http.Handler {
	return middleware(routes())
}

// LoadHTTP returns the HTTPS routes and middleware
func LoadHTTP() http.Handler {
	return middleware(routes())

	// Uncomment this and comment out the line above to always redirect to HTTPS
	//return http.HandlerFunc(redirectToHTTPS)
}

// Optional method to make it easy to redirect from HTTP to HTTPS
func redirectToHTTPS(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host, http.StatusMovedPermanently)
}

// *****************************************************************************
// Routes
// *****************************************************************************

func routes() *httprouter.Router {
	r := httprouter.New()

	// Get WebSocket
	r.GET("/:id", Handler(alice.
		New().
		ThenFunc(ConnectWebsocket)))

	return r
}

func ConnectWebsocket(w http.ResponseWriter, r *http.Request) {
	Melody.HandleRequest(w, r)
}

// *****************************************************************************
// Middleware
// *****************************************************************************

func middleware(h http.Handler) http.Handler {

	// Log every request
	h = HandleLogs(h)

	return h
}

// Handler will log the HTTP requests
func HandleLogs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// HandlerFunc accepts the name of a function so you don't have to wrap it with http.HandlerFunc
// Example: r.GET("/", httprouterwrapper.HandlerFunc(controller.Index))
func HandlerFunc(h http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}

// Handler accepts a handler to make it compatible with http.HandlerFunc
// Example: r.GET("/", httprouterwrapper.Handler(http.HandlerFunc(controller.Index)))
func Handler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}
