package router

import (
	"net/http"
)

type Router struct {
	routes       []*RouteConfig
	staticRoutes []*StaticRouteConfig
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
		methods:             httpMethods[:],
		acceptTrailingSlash: true,
		csrfProtect:         false,
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

// Serve static files on the configured path
func (router *Router) ServeStatic(urlpath string, dirpath string) *StaticRouteConfig {
	rc := StaticRouteConfig{
		route:     urlpath,
		dirpath:   dirpath,
		indexFile: "",
	}
	router.staticRoutes = append(router.staticRoutes, &rc)
	return &rc
}
