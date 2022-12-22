package routex

import (
	"net/http"
	"strings"
)

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, rt := range router.routes {

		// MATCHING THE ROUTE
		//check of accept trailing-slash settings. By default the router will accept them even if none is set in the Handle-Path
		rt.route = strings.TrimRight(rt.route, "/")
		if rt.acceptTrailingSlash && rt.route+"/" == r.URL.Path || rt.route == r.URL.Path {

			// OTHER CHECKS

			//check allowed methods
			if !rt.checkMethods(r) {
				forbidden(w)
				return
			}

			//check if CORS are allowed, if so, set headers
			rt.checkAllowCors(w)

			// finally execute the configured HandlerFunc
			rt.handlerfn(w, r)
			return
		}
	}
	// STATIC SERVING (will be always overruled in case of double configuration)
	// check if route is configured for serving static files and serve them
	for _, rt := range router.staticRoutes {
		rt.route = strings.TrimRight(rt.route, "/")
		if strings.Index(r.URL.Path, rt.route) == 0 {
			staticServing(r.URL.Path, rt, w)
			return
		}
	}

	//NOT FOUND
	//if the function rund until here, no route was configured for this request. Return a not found error.
	notfound(w)
}

// error funcs
func forbidden(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
}

func notfound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}
