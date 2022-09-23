package router

import (
	"net/http"
	"strings"
)

type router struct {
	routes []*routeConfig
}

// Create a new Router
func New() router {
	return router{}
}

// Connect an URL-Path to a handler function.
func (router *router) Handle(path string, handlerfn http.HandlerFunc) *routeConfig {
	rc := routeConfig{
		route:               path,
		handlerfn:           handlerfn,
		methods:             []string{},
		acceptTrailingSlash: true,
		csrfProtect:         false,
		serveStatic:         false,
		staticPath:          "",
		indexFile:           "",
	}
	router.routes = append(router.routes, &rc)
	return &rc
}

// Serve static files on the configured path
func (router *router) ServeStatic(urlpath string, dirpath string) *routeConfig {
	rc := routeConfig{
		route:               urlpath,
		methods:             []string{},
		acceptTrailingSlash: true,
		csrfProtect:         false,
		serveStatic:         true,
		staticPath:          dirpath,
		indexFile:           "",
	}
	router.routes = append(router.routes, &rc)
	return &rc
}

// Accept a trailing slash for all configured routes if client added one to the request even if you configured none in your path.
func (router *router) AcceptTrailingSlash(b bool) {
	for _, rt := range router.routes {
		rt.acceptTrailingSlash = b
	}
}

// Set allowed methods for all configured routes.
func (router *router) Methods(methods ...string) {
	for _, rt := range router.routes {
		rt.methods = methods
	}
}

func (router router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, rt := range router.routes {

		//check allowed methods
		if !rt.checkMethods(r) {
			forbidden(w, r)
			return
		}

		//check if static serving is enabled
		if rt.serveStatic {
			staticServing(w, r, rt)
			return
		}

		//check and handle trailing slash configs
		rt.route = strings.TrimRight(rt.route, "/")
		if rt.acceptTrailingSlash && rt.route+"/" == r.URL.Path {
			rt.handlerfn(w, r)
		} else if rt.route == r.URL.Path {
			rt.handlerfn(w, r)
		}
	}
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
}
