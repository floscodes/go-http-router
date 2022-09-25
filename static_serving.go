package router

import (
	"net/http"
	"os"
	"strings"
)

func staticServing(path string, rt *StaticRouteConfig, w http.ResponseWriter) {
	// making urlpath ready to be partially substituted by the root of the file directory
	path = strings.TrimRight(path, "/")
	path = path + "/"
	path = strings.TrimLeft(path, "/")
	if strings.Contains(path, "/") {
		//substitution
		path = path[strings.Index(path, "/")+1:]
	}
	// preparing file directory path
	rt.dirpath = strings.TrimRight(rt.dirpath, "/")
	// building file path to open
	path = rt.dirpath + "/" + path
	//removing trailing slash for opening and serving a static file
	path = strings.TrimRight(path, "/")

	if file, err := os.ReadFile(path); err != nil {
		if rt.indexFile == "" {
			if file, err = os.ReadFile(path + "/" + "index.html"); err != nil {
				w.Write([]byte(""))
			} else {
				w.Write(file)
			}
		} else {
			if file, err = os.ReadFile(path + "/" + rt.indexFile); err != nil {
				w.Write([]byte(""))
			} else {
				w.Write(file)
			}
		}
	} else {
		w.Write(file)
	}
}
