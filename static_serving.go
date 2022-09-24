package router

import (
	"net/http"
	"os"
	"strings"
)

func (rt *RouteConfig) staticServing(w http.ResponseWriter, r *http.Request) {
	/*if strings.TrimRight(r.URL.Path, "/") != strings.TrimRight(rt.staticPath, "/") {
		return
	}*/
	path := r.URL.Path[strings.Index(r.URL.Path, rt.route)+len(rt.route):]
	path = strings.TrimLeft(path, "/")
	path = "/" + path
	if path == "/" {
		if rt.indexFile == "" {
			path = path + "/index.html"
		} else {
			path = path + "/" + rt.indexFile
		}
	}
	rt.staticPath = strings.TrimRight(rt.staticPath, "/")
	path = rt.staticPath + path
	if file, err := os.ReadFile(path); err != nil {
		w.Write([]byte(""))
	} else {
		w.Write(file)
	}
}
