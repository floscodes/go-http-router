package router

type StaticRouteConfig struct {
	route     string
	dirpath   string
	indexFile string
}

// Set custom index file when serving static files.
func (rc *StaticRouteConfig) IndexFile(file string) {
	rc.indexFile = file
}
