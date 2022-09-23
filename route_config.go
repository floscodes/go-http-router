package router

import (
	"net/http"
)

type routeConfig struct {
	route               string
	handlerfn           http.HandlerFunc
	methods             []string
	acceptTrailingSlash bool
	csrfProtect         bool
	serveStatic         bool
	staticPath          string
	indexFile           string
}

// Accept a trailing slash a single route if client added one to the request even if you configured none in your path.
func (rc *routeConfig) AcceptTrailingSlash(b bool) {
	rc.acceptTrailingSlash = b
}

// Set allowed methods for a single route.
func (rc *routeConfig) Methods(methods ...string) {
	rc.methods = append(rc.methods, methods...)
}

// Set custom index file when serving static files.
func (rc *routeConfig) IndexFile(file string) {
	rc.indexFile = file
}
