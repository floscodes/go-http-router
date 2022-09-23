<table><tr>
<td><img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" width="70" height="60"></td><td><h1>Easy to use router for http-services</h1></td>
	<td>
	<h3>Set routes and link them to functions</h3>
	</td>
<br>
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
```
router.ServeStatic("/static", "./static").IndexFile("template.html")
```
