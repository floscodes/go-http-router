package router

import (
	"net/http"
	"strings"
)

func (rt *routeConfig) checkMethods(r *http.Request) bool {
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
