[![Go Reference](https://pkg.go.dev/badge/github.com/floscodes/go-http-router.svg)](https://pkg.go.dev/github.com/floscodes/go-http-router)
# Easy to use router for http-services 

### Set routes and link them to functions

```go
func main() {
	router := router.New()

	router.Handle("/hello", hello)

	fmt.Println("Server is listening")
	http.ListenAndServe(":8080", router)

}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello!"))
}
```

You can specify the allowed methods like this:
```go
router.Handle("/hello", hello).Methods("POST", "GET")
```

Or set them for all configured paths:
```go
router.Methods("POST", "GET")
```

The router will accept trailing slashes typed by the client by default, even if none is set in the handling path.
You can disable that this way for a certain route...
```go
router.Handle("/hello", hello).AcceptTrailingSlash(false)
```
... or disable it for all routes:
```go
router.AcceptTrailingSlash(false)
```

### Serve static files

```go
router.ServeStatic("/static", "./static")
```

If the request does not contain a specific filename, the router will automatically look for `index.html`.

Optionally you can set a custom index file:
```go
router.ServeStatic("/static", "./static").IndexFile("template.html")
```

### Allow CORS

You can just allow CORS like this:
```go
router.AllowCORS(true)
```

or for a single route.
```go
router.Handle("/api", api).AllowCORS(true)
```

**Whenever allowing CORS it is highly recommended to define the allowed methods, otherwise all methods will be accepted!**

However, custom CORS heraders could be set by yourself at anytime in your handler functions.

```go
router.Handle("/api", api).AllowCORS(true).Methods("GET")
```
