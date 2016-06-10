package route

import (
	"net/http"

	"github.com/evanfeenstra/circuitSocket/app/controller"
	hr "github.com/evanfeenstra/circuitSocket/app/route/middleware/httprouterwrapper"
	"github.com/evanfeenstra/circuitSocket/app/route/middleware/logrequest"

	"github.com/gorilla/context"
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

	// Set 404 handler
	/*r.NotFound = alice.
		New().
		ThenFunc(controller.Error404)*/

	// Get WebSocket
	r.GET("/:id/ws", hr.Handler(alice.
		New().
		ThenFunc(controller.ReadWS)))
	

	return r
}

// *****************************************************************************
// Middleware
// *****************************************************************************

func middleware(h http.Handler) http.Handler {

	// Log every request
	h = logrequest.Handler(h)

	// Clear handler for Gorilla Context (without this, memory leak)
	h = context.ClearHandler(h)

	return h
}
