package router

import (
	"net/http"
	"strings"
)

var httpMethods = [9]string{"GET", "POST", "PUT", "DELETE", "HEAD", "CONNECT", "OPTIONS", "TRACE", "PATCH"}

type RouteConfig struct {
	route               string
	handlerfn           http.HandlerFunc
	methods             []string
	acceptTrailingSlash bool
	allowCors           bool
}

// Accept a trailing slash a single route if client added one to the request even if you configured none in your path.
func (rc *RouteConfig) AcceptTrailingSlash(b bool) {
	rc.acceptTrailingSlash = b
}

// Set allowed methods for a single route.
func (rc *RouteConfig) Methods(methods ...string) {
	rc.methods = methods
}

// Set headers to allow CORS requests for a single route. Default is false.
func (rc *RouteConfig) AllowCORS(b bool) {
	rc.allowCors = b
}

// checking functions - not public

func (rt *RouteConfig) checkMethods(r *http.Request) bool {
	if len(rt.methods) == 0 {
		return true
	}
	for _, m := range rt.methods {
		if strings.EqualFold(m, r.Method) {
			return true
		}
	}
	return false
}

func (rt *RouteConfig) checkAllowCors(w http.ResponseWriter) {
	if !rt.allowCors {
		return
	} else {
		if len(rt.methods) == 0 {
			rt.methods = httpMethods[:]
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(rt.methods, ", "))
		w.Header().Set("Accept", "*/*")
	}
}
