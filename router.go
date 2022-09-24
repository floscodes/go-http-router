package router

import (
	"net/http"
)

type Router struct {
	routes []*RouteConfig
}

// Create a new Router
func New() Router {
	return Router{}
}

// Connect an URL-Path to a handler function.
func (router *Router) Handle(path string, handlerfn http.HandlerFunc) *RouteConfig {
	rc := RouteConfig{
		route:               path,
		handlerfn:           handlerfn,
		methods:             []string{},
		acceptTrailingSlash: true,
		csrfProtect:         false,
		serveStatic:         false,
		staticPath:          "",
		indexFile:           "",
		allowCors:           false,
	}
	router.routes = append(router.routes, &rc)
	return &rc
}

// Serve static files on the configured path
func (router *Router) ServeStatic(urlpath string, dirpath string) *RouteConfig {
	rc := RouteConfig{
		route:               urlpath,
		methods:             []string{},
		acceptTrailingSlash: true,
		csrfProtect:         false,
		serveStatic:         true,
		staticPath:          dirpath,
		indexFile:           "",
		allowCors:           false,
	}
	router.routes = append(router.routes, &rc)
	return &rc
}

// Accept a trailing slash for all configured routes if client added one to the request even if you configured none in your path.
func (router *Router) AcceptTrailingSlash(b bool) {
	for _, rt := range router.routes {
		rt.acceptTrailingSlash = b
	}
}

// Set allowed methods for all configured routes.
func (router *Router) Methods(methods ...string) {
	for _, rt := range router.routes {
		rt.methods = methods
	}
}

// Set headers to allow CORS requests for all configured routes. Default is false.
func (router *Router) AllowCORS(b bool) {
	for _, rt := range router.routes {
		rt.allowCors = b
	}
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
}
