package router

import (
	"fmt"
	"net/http"
	"strings"
)

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, rt := range router.routes {

		// MATCHING THE ROUTE
		//check of accept trailing-slash settings. By default the router will accept them even if none is set in the Handle-Path
		rt.route = strings.TrimRight(rt.route, "/")
		if rt.acceptTrailingSlash && rt.route+"/" == r.URL.Path || rt.route == r.URL.Path {
			fmt.Println("requested")
			// OTHER CHECKS

			//check allowed methods
			if !rt.checkMethods(r) {
				forbidden(w, r)
				return
			}

			//check if CORS are allowed, if so, set headers
			rt.checkAllowCors(w)

			//check if static serving is enabled
			if rt.serveStatic {
				rt.staticServing(w, r)
				return
			}

			// finally execute the configured HandlerFunc
			rt.handlerfn(w, r)
		}
	}
}
