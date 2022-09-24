package router

import (
	"net/http"
	"strings"
)

var httpMethods = [9]string{"GET", "POST", "PUT", "DELETE", "HEAD", "CONNECT", "OPTIONS", "TRACE", "PATCH"}

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
