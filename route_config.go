package router

import (
	"net/http"
)

type RouteConfig struct {
	route               string
	handlerfn           http.HandlerFunc
	methods             []string
	acceptTrailingSlash bool
	csrfProtect         bool
	serveStatic         bool
	staticPath          string
	indexFile           string
	allowCors           bool
}

// Accept a trailing slash a single route if client added one to the request even if you configured none in your path.
func (rc *RouteConfig) AcceptTrailingSlash(b bool) {
	rc.acceptTrailingSlash = b
}

// Set allowed methods for a single route.
func (rc *RouteConfig) Methods(methods ...string) {
	rc.methods = append(rc.methods, methods...)
}

// Set custom index file when serving static files.
func (rc *RouteConfig) IndexFile(file string) {
	rc.indexFile = file
}

// Set headers to allow CORS requests for a single route. Default is false.
func (rc *RouteConfig) AllowCORS(b bool) {
	rc.allowCors = b
}
